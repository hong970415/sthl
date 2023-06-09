import { useState } from 'react'
import { useFormik } from 'formik'
import { get } from 'lodash'
import {
  Button,
  Flex,
  NumberInput,
  Select,
  Skeleton,
  Text,
  Textarea,
  TextInput,
  useMantineTheme,
} from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import { GetProductStatusListWithLabel } from '@/constants/constants'
import { API, StatusCode } from '@/services'
import { IProduct, ProductFieldKeys } from '@/entities/product'
import useGetProductById from '@/hooks/api/useGetProductById'
import useTranslationData from '@/hooks/useTranslation'
import useAuth from '@/hooks/useAuth'
import { showDate } from '@/utils/date'
import { ProductFormSchema } from './ProductCreateForm'
import ProductStatusBadge from '../badges/ProductStatusBadge'
import SelectImg from '../SelectImg/SelectImg'
import useGetAlbumImgsByUserId from '@/hooks/api/useGetAlbumImgsByUserId'
import AlbumImg from '../AlbumList/AlbumImg'

function getEditableProductField(product: IProduct) {
  return {
    name: product.name,
    price: product.price,
    quantity: product.quantity,
    description: product.description,
    status: product.status,
    imgUrl: product.imgUrl,
  }
}
export interface IProductEditFormProps {
  productId: string
}
export default function ProductEditForm(props: IProductEditFormProps) {
  const { productId } = props
  const theme = useMantineTheme()
  const { t } = useTranslationData()
  const { album, isFetched } = useGetAlbumImgsByUserId({ limit: 1000 })
  const { user } = useAuth()
  const userId = user ? user.id : ''
  const { product, isLoading, refetch } = useGetProductById(userId, productId)

  const [state, setState] = useState({ isEditing: false })
  const editProductForm = useFormik({
    initialValues: {
      name: '',
      price: 0,
      quantity: 0,
      description: '',
      status: '',
      imgUrl: '',
    },
    validationSchema: ProductFormSchema,
    onSubmit: async (values) => {
      // console.log('ProductEditForm values', values)
      const payload = {
        name: values.name,
        price: values.price,
        quantity: values.quantity,
        description: values.description,
        status: values.status,
        imgUrl: values.imgUrl,
      }
      const response = await API.putUpdateProductById(
        userId,
        productId,
        payload
      )
      if (response.success && response.status === StatusCode.Ok) {
        showNotification({ color: 'green', message: response.data.msg })
        refetch()
        setState((prev) => ({ ...prev, isEditing: false }))
      } else {
        const message = response.errorMsg
        showNotification({ color: 'red', message: message })
      }
    },
  })

  // actions
  const handleOnClickEdit = (event: React.MouseEvent<HTMLButtonElement>) => {
    if (product) {
      editProductForm.resetForm({
        values: getEditableProductField(product),
      })
      setState((prev) => ({ ...prev, isEditing: true }))
    }
  }

  const handleOnClickCancel = (event: React.MouseEvent<HTMLButtonElement>) => {
    if (product) {
      setState((prev) => ({ ...prev, isEditing: false }))
      editProductForm.resetForm({
        values: getEditableProductField(product),
      })
    }
  }

  // ui elements
  const viewModeEl = (
    <Flex direction={'column'}>
      <Skeleton visible={isLoading}>
        <Flex align={'center'} mb="md">
          <Text size="lg" weight="bold" sx={{ width: '20%' }}>
            {t('product.img')}:
          </Text>
          <div style={{ width: '240px' }}>
            <AlbumImg imgSrc={product ? product.imgUrl : ''} />
          </div>
        </Flex>
      </Skeleton>
      {ProductFieldKeys.map((key) => {
        if (key === 'createdAt' || key === 'updatedAt') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text size="lg" weight="bold" sx={{ width: '20%' }}>
                  {t(`product.${key}`)}:
                </Text>
                <Text sx={{ width: '80%' }}>
                  {showDate(get(product, key, ''))}
                </Text>
              </Flex>
            </Skeleton>
          )
        }
        if (key === 'status') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex key={key} align={'center'} mb="md">
                <Text size="lg" weight="bold" sx={{ width: '20%' }}>
                  {t(`product.${key}`)}:
                </Text>
                <Text sx={{ width: '80%' }}>
                  {product?.status && (
                    <ProductStatusBadge value={product.status} />
                  )}
                </Text>
              </Flex>
            </Skeleton>
          )
        }
        return (
          <Skeleton key={key} visible={isLoading}>
            <Flex key={key} align={'center'} mb="md">
              <Text size="lg" weight="bold" sx={{ width: '20%' }}>
                {t(`product.${key}`)}:
              </Text>
              <Text sx={{ width: '80%' }}>{get(product, key, '')}</Text>
            </Flex>
          </Skeleton>
        )
      })}
      <Flex justify={'center'}>
        <Button
          mt="md"
          mb="md"
          type="submit"
          sx={{ fontSize: theme.fontSizes.md }}
          onClick={handleOnClickEdit}
          disabled={isLoading}
        >
          {t('general.edit')}
        </Button>
      </Flex>
    </Flex>
  )

  const editModeEl = (
    <form onSubmit={editProductForm.handleSubmit}>
      <SelectImg
        name={'imgUrl'}
        label={t('product.img')}
        data={
          album && album.imgs
            ? album.imgs.map((el) => ({
                value: el.imgUrl,
                label: el.imgName,
                imgUrl: el.imgUrl,
              }))
            : []
        }
        value={editProductForm.values.imgUrl}
        onChange={(value) => {
          editProductForm.setFieldValue('imgUrl', value)
        }}
        disabled={editProductForm.isSubmitting}
      />
      <TextInput
        my="md"
        name="name"
        label={t('product.name')}
        value={editProductForm.values.name}
        onChange={editProductForm.handleChange}
        error={
          editProductForm.errors.name && editProductForm.touched.name
            ? t(`error.${editProductForm.errors.name}`)
            : null
        }
        disabled={editProductForm.isSubmitting}
        withAsterisk
      />
      <NumberInput
        mb="md"
        name="price"
        label={t('product.price')}
        type="number"
        min={0}
        step={0.01}
        precision={2}
        value={editProductForm.values.price}
        onChange={(value) => {
          editProductForm.setFieldValue('price', value)
        }}
        error={
          editProductForm.errors.price && editProductForm.touched.price
            ? t(`error.${editProductForm.errors.price}`)
            : null
        }
        disabled={editProductForm.isSubmitting}
        withAsterisk
      />
      <NumberInput
        mb="md"
        name="quantity"
        label={t('product.quantity')}
        type="number"
        min={1}
        value={editProductForm.values.quantity}
        onChange={(value) => {
          editProductForm.setFieldValue('quantity', value)
        }}
        error={
          editProductForm.errors.quantity && editProductForm.touched.quantity
            ? t(`error.${editProductForm.errors.quantity}`)
            : null
        }
        disabled={editProductForm.isSubmitting}
        withAsterisk
      />
      <Textarea
        mb="md"
        name="description"
        label={t('product.description')}
        value={editProductForm.values.description}
        onChange={editProductForm.handleChange}
        disabled={editProductForm.isSubmitting}
        error={
          editProductForm.errors.description &&
          editProductForm.touched.description
            ? t(`error.${editProductForm.errors.description}`)
            : null
        }
      />
      <Select
        mb="md"
        name="status"
        label={t('product.status')}
        data={GetProductStatusListWithLabel(t)}
        value={editProductForm.values.status}
        onChange={(value) => {
          editProductForm.setFieldValue('status', value)
        }}
        disabled={editProductForm.isSubmitting}
      />
      <Flex justify="center" align="center" gap={'lg'}>
        <Button
          mt="md"
          mb="md"
          type="submit"
          sx={{ fontSize: theme.fontSizes.md }}
          loading={editProductForm.isSubmitting}
        >
          {t('general.update')}
        </Button>
        <Button
          mt="md"
          mb="md"
          type="submit"
          variant="outline"
          sx={{ fontSize: theme.fontSizes.md }}
          color="gray"
          onClick={handleOnClickCancel}
          disabled={editProductForm.isSubmitting}
        >
          {t('general.cancel')}
        </Button>
      </Flex>
    </form>
  )
  return state.isEditing ? editModeEl : viewModeEl
}
