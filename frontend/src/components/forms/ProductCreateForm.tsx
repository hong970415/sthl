import { forwardRef } from 'react'
import { useFormik } from 'formik'
import * as yup from 'yup'
import {
  Button,
  Flex,
  NumberInput,
  Textarea,
  TextInput,
  useMantineTheme,
} from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import { API, ICreateProductForm, StatusCode } from '@/services'
import useTranslationData from '@/hooks/useTranslation'
import useGetAlbumImgsByUserId from '@/hooks/api/useGetAlbumImgsByUserId'
import SelectImg from '../SelectImg/SelectImg'

export const ProductFormSchema = yup.object().shape({
  name: yup
    .string()
    .required('required_error')
    .max(50, 'product_form_name_length_error'),
  price: yup
    .number()
    .required('required_error')
    .moreThan(0, 'product_form_price_moreThan_error'),
  quantity: yup
    .number()
    .required('required_error')
    .min(0, 'product_form_quantity_min_error'),
  description: yup.string().max(200, 'product_form_description_length_error'),
  imgUrl: yup.string().max(512, ''),
})

export default function ProductCreateForm() {
  const theme = useMantineTheme()
  const { t } = useTranslationData()
  const { album, isFetched } = useGetAlbumImgsByUserId({ limit: 1000 })
  const createProductForm = useFormik<ICreateProductForm>({
    initialValues: {
      name: '',
      price: 0,
      quantity: 0,
      description: '',
      imgUrl: '',
    },
    validationSchema: ProductFormSchema,
    onSubmit: async (values, { resetForm }) => {
      // console.log('ProductCreateForm values', values)
      const payload = {
        name: values.name,
        price: values.price,
        quantity: values.quantity,
        description: values.description,
        imgUrl: values.imgUrl,
      }
      const response = await API.postCreateProduct(payload)
      if (response.success && response.status === StatusCode.Created) {
        const message = response.data.msg
        showNotification({ color: 'green', message: message })
        resetForm()
      } else {
        const message = response.errorMsg
        showNotification({ color: 'red', message: message })
      }
    },
  })

  // actions
  const handleOnClickClear = (event: React.MouseEvent<HTMLButtonElement>) => {
    createProductForm.resetForm()
  }
  return (
    <form onSubmit={createProductForm.handleSubmit}>
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
        value={createProductForm.values.imgUrl}
        onChange={(value) => {
          createProductForm.setFieldValue('imgUrl', value)
        }}
        disabled={createProductForm.isSubmitting}
      />
      <TextInput
        my="md"
        name="name"
        label={t('product.name')}
        value={createProductForm.values.name}
        onChange={createProductForm.handleChange}
        error={
          createProductForm.errors.name && createProductForm.touched.name
            ? t(`error.${createProductForm.errors.name}`)
            : null
        }
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
        value={createProductForm.values.price}
        onChange={(value) => {
          createProductForm.setFieldValue('price', value)
        }}
        error={
          createProductForm.errors.price && createProductForm.touched.price
            ? t(`error.${createProductForm.errors.price}`)
            : null
        }
        withAsterisk
      />
      <NumberInput
        mb="md"
        name="quantity"
        label={t('product.quantity')}
        type="number"
        min={1}
        value={createProductForm.values.quantity}
        onChange={(value) => {
          createProductForm.setFieldValue('quantity', value)
        }}
        error={
          createProductForm.errors.quantity &&
          createProductForm.touched.quantity
            ? t(`error.${createProductForm.errors.quantity}`)
            : null
        }
        withAsterisk
      />
      <Textarea
        name="description"
        label={t('product.description')}
        value={createProductForm.values.description}
        onChange={createProductForm.handleChange}
        error={
          createProductForm.errors.description &&
          createProductForm.touched.description
            ? t(`error.${createProductForm.errors.description}`)
            : null
        }
      />
      <Flex justify="center" align="center" gap={'lg'}>
        <Button
          mt="md"
          mb="md"
          type="submit"
          sx={{ fontSize: theme.fontSizes.md }}
          loading={createProductForm.isSubmitting}
        >
          {t('general.create')}
        </Button>
        <Button
          mt="md"
          mb="md"
          variant="outline"
          sx={{ fontSize: theme.fontSizes.md }}
          color={'gray'}
          onClick={handleOnClickClear}
        >
          {t('general.clear')}
        </Button>
      </Flex>
    </form>
  )
}
