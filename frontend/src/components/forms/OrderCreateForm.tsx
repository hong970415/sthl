import {
  ActionIcon,
  Button,
  Center,
  Flex,
  NumberInput,
  Select,
  Text,
  Textarea,
  TextInput,
  useMantineTheme,
} from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import { useFormik } from 'formik'
import * as yup from 'yup'
import { get } from 'lodash'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faMinus, faPlus } from '@fortawesome/free-solid-svg-icons'
import { API, ICreateOrderForm, ICreateOrderItem, StatusCode } from '@/services'
import useTranslationData from '@/hooks/useTranslation'
import useGetProducts from '@/hooks/api/useGetProducts'
import { GetPaymentMethodListWithLabel } from '@/constants/constants'
import {
  calculateOrderTotalAmount,
  getFormattedOrderTotalAmount,
} from '@/utils/order'
import useAuth from '@/hooks/useAuth'

export const CreateOrderFormSchema = yup.object().shape({
  items: yup
    .array()
    .of(
      yup.object({
        productId: yup
          .string()
          .required('required_error')
          .uuid('required_error'),
        quantity: yup
          .number()
          .required('required_error')
          .min(1, 'order_form_quantity_min_error'),
        purchasedPrice: yup
          .number()
          .required('required_error')
          .moreThan(0, 'order_form_purchasedPrice_moreThan_error'),
      })
    )
    .required('required_error'),
  remark: yup.string().max(255, 'order_form_remark_length_error'),
  discount: yup.number().min(0.4, 'required_error').max(1, 'required_error'),
  // totalAmount: yup.string().max(200, 'order_form_remark_length_error'),
  paymentMethod: yup
    .string()
    .required('required_error')
    .max(255, 'required_error'),
  shippingAddress: yup
    .string()
    .required('required_error')
    .max(512, 'required_error'),
})
export function getOrderItemInitialValue() {
  return {
    id: '',
    productId: '',
    quantity: 0,
    purchasedPrice: 0,
    purchasedName: '',
  }
}
export default function OrderCreateForm() {
  const theme = useMantineTheme()
  const { t } = useTranslationData()
  const hasAdminPermission = false
  const { products, refetch } = useGetProducts({ limit: 1000 })
  const { user } = useAuth()
  const createOrderForm = useFormik<ICreateOrderForm>({
    initialValues: {
      items: [getOrderItemInitialValue()],
      remark: '',
      discount: 1,
      totalAmount: 0,
      paymentMethod: '',
      shippingAddress: '',
    },
    validationSchema: CreateOrderFormSchema,
    onSubmit: async (values, { resetForm }) => {
      // console.log('OrderCreateForm values', values)
      if (!user) {
        return
      }
      const payload = {
        items: values.items,
        remark: values.remark,
        discount: values.discount,
        totalAmount: calculateOrderTotalAmount(values.items, values.discount),
        paymentMethod: values.paymentMethod,
        shippingAddress: values.shippingAddress,
      }
      const response = await API.postCreateOrder(user.id, payload)
      if (response.success && response.status === StatusCode.Created) {
        const message = response.data.msg
        refetch()
        resetForm()
        showNotification({ color: 'green', message: message })
      } else {
        const message = response.errorMsg
        showNotification({ color: 'red', message: message })
      }
    },
  })

  // helpers
  const findProductById = (productId: string) => {
    return products.find((item) => item.id === productId)
  }
  // get available product quantity after chosen product
  const getAvailableProductQuantity = (productItem: ICreateOrderItem) => {
    const flattedOrderItemByIdExcludeSelf = createOrderForm.values.items
      .map((item) => item.productId)
      .filter((item) => item !== productItem.productId)
    const availableProducts = products
      .filter((item) => !flattedOrderItemByIdExcludeSelf.includes(item.id))
      .map((item) => ({ value: item.id, label: item.name }))
    const maxProductQuantity = findProductById(productItem.productId)?.quantity
    let productItemQuantity = productItem.quantity || 0
    let availableProductQuantity = maxProductQuantity
    if (availableProductQuantity) {
      availableProductQuantity -= productItemQuantity
    }
    return { availableProducts, availableProductQuantity, maxProductQuantity }
  }

  // forms
  const handleAddItem = () => {
    createOrderForm.setFieldValue('items', [
      ...createOrderForm.values.items,
      getOrderItemInitialValue(),
    ])
  }

  const handleOnChangeItem = (productId: string, index: number) => {
    const selectedProduct = findProductById(productId)
    const defaultValue = {
      productId: productId,
      quantity: 1,
      purchasedPrice: selectedProduct?.price,
      purchasedName: selectedProduct?.name,
    }
    createOrderForm.setFieldValue(`items[${index}]`, defaultValue)
  }

  const handleOnChangeItemQuantity = (index: number, value: number) => {
    createOrderForm.setFieldValue(`items[${index}].quantity`, value)
  }

  const handleOnChangeItemPrice = (index: number, value: number) => {
    createOrderForm.setFieldValue(`items[${index}].purchasedPrice`, value)
  }

  const handleOnChangeDiscount = (value: number) => {
    createOrderForm.setFieldValue('discount', value)
  }
  const handleOnChangePaymentMethod = (value: string) => {
    createOrderForm.setFieldValue('paymentMethod', value)
  }

  const handleRemoveItem = (productId: string) => {
    createOrderForm.setFieldValue(
      'items',
      createOrderForm.values.items.filter(
        (item) => item.productId !== productId
      )
    )
  }

  const handleOnClickClear = (event: React.MouseEvent<HTMLButtonElement>) => {
    createOrderForm.resetForm()
  }

  // ui elements
  const addItemButton = (
    <Center>
      <ActionIcon
        variant="outline"
        size="lg"
        radius={'lg'}
        onClick={() => {
          handleAddItem()
        }}
      >
        <FontAwesomeIcon icon={faPlus} />
      </ActionIcon>
    </Center>
  )
  const removeItemButton = (value: string, index: number) => (
    <ActionIcon
      variant="outline"
      size="lg"
      radius={'lg'}
      onClick={() => {
        handleRemoveItem(value)
      }}
    >
      <FontAwesomeIcon icon={faMinus} />
    </ActionIcon>
  )
  // form elements
  const formItemsEl =
    createOrderForm.values.items.length > 0
      ? createOrderForm.values.items.map((productItem, index) => {
          const {
            availableProducts,
            availableProductQuantity,
            maxProductQuantity,
          } = getAvailableProductQuantity(productItem)
          return (
            <Flex mb="md" key={productItem.productId} gap="md" align={'end'}>
              <Select
                name={`items[${index}].productId`}
                label={t('order.items')}
                data={availableProducts}
                value={productItem.productId}
                onChange={(value) => handleOnChangeItem(value as string, index)}
                error={
                  get(
                    createOrderForm,
                    `errors.items[${index}].productId`,
                    null
                  ) &&
                  get(
                    createOrderForm,
                    `touched.items[${index}].productId`,
                    null
                  )
                    ? t(
                        `error.${get(
                          createOrderForm,
                          `errors.items[${index}].productId`,
                          ''
                        )}`
                      )
                    : null
                }
                sx={{ width: '100%' }}
                withAsterisk
              />
              <NumberInput
                name={`items[${index}].quantity`}
                label={t('order.quantity')}
                description={
                  availableProductQuantity &&
                  `${t('order.avail')} ${availableProductQuantity}`
                }
                type="number"
                min={1}
                max={maxProductQuantity}
                value={productItem.quantity}
                onChange={(value) => {
                  handleOnChangeItemQuantity(index, value as number)
                }}
                error={
                  get(
                    createOrderForm,
                    `errors.items[${index}].quantity`,
                    null
                  ) && createOrderForm.submitCount > 0
                    ? t(
                        `error.${get(
                          createOrderForm,
                          `errors.items[${index}].quantity`,
                          ''
                        )}`
                      )
                    : null
                }
                sx={{ width: '100%' }}
                withAsterisk
              />
              <NumberInput
                name={`items[${index}].purchasedPrice`}
                label={t('order.purchasedPrice')}
                type="number"
                min={0}
                step={0.01}
                precision={2}
                value={productItem.purchasedPrice}
                onChange={(value) => {
                  handleOnChangeItemPrice(index, value as number)
                }}
                error={
                  get(
                    createOrderForm,
                    `errors.items[${index}].purchasedPrice`,
                    null
                  ) &&
                  get(
                    createOrderForm,
                    `touched.items[${index}].purchasedPrice`,
                    null
                  )
                    ? t(
                        `error.${get(
                          createOrderForm,
                          `errors.items[${index}].purchasedPrice`,
                          ''
                        )}`
                      )
                    : null
                }
                disabled={!hasAdminPermission}
                sx={{ width: '100%' }}
                withAsterisk
              />
              {createOrderForm.values.items.length > 1
                ? removeItemButton(productItem.productId, index)
                : null}
            </Flex>
          )
        })
      : null

  const calculateEl = (
    <Flex justify={'flex-end'} mr="lg">
      <Flex align={'center'}>
        <Text size={'lg'}>{t('order.totalAmount')}: </Text>
        <Text>
          {getFormattedOrderTotalAmount(
            createOrderForm.values.items,
            createOrderForm.values.discount
          )}
        </Text>
      </Flex>
    </Flex>
  )

  const formDiscountEl = (
    <Flex justify={'flex-end'} mb="md">
      <NumberInput
        name={'discount'}
        label={t('order.discount')}
        min={0.4}
        max={1}
        precision={2}
        step={0.01}
        value={createOrderForm.values.discount}
        onChange={handleOnChangeDiscount}
        error={
          createOrderForm.errors.discount && createOrderForm.touched.discount
            ? t(`error.${createOrderForm.errors.discount}`)
            : null
        }
        sx={{ width: '100%' }}
        withAsterisk
      />
    </Flex>
  )

  const formPaymentMethodEl = (
    <Select
      name={'paymentMethod'}
      label={t('order.paymentMethod')}
      data={GetPaymentMethodListWithLabel(t)}
      value={createOrderForm.values.paymentMethod}
      onChange={handleOnChangePaymentMethod}
      error={
        createOrderForm.errors.paymentMethod &&
        createOrderForm.touched.paymentMethod
          ? t(`error.${createOrderForm.errors.paymentMethod}`)
          : null
      }
      withAsterisk
      mt="md"
      mb="md"
    />
  )
  const formShippingAddressEl = (
    <Textarea
      mb={'md'}
      name="shippingAddress"
      label={t('order.shippingAddress')}
      value={createOrderForm.values.shippingAddress}
      onChange={createOrderForm.handleChange}
      error={
        createOrderForm.errors.shippingAddress &&
        createOrderForm.touched.shippingAddress
          ? t(`error.${createOrderForm.errors.shippingAddress}`)
          : null
      }
      withAsterisk
    />
  )
  const formRemarkEl = (
    <Textarea
      name="remark"
      label={t('order.remark')}
      value={createOrderForm.values.remark}
      onChange={createOrderForm.handleChange}
      error={
        createOrderForm.errors.remark && createOrderForm.touched.remark
          ? t(`error.${createOrderForm.errors.remark}`)
          : null
      }
    />
  )
  // console.log('createOrderForm', createOrderForm)
  return (
    <form onSubmit={createOrderForm.handleSubmit}>
      {formItemsEl}
      {createOrderForm.values.items.length !== products.length && addItemButton}
      {formDiscountEl}
      {calculateEl}
      {formPaymentMethodEl}
      {formShippingAddressEl}
      {formRemarkEl}
      <Flex justify="center" align="center" gap={'lg'}>
        <Button
          mt="md"
          mb="md"
          type="submit"
          sx={{ fontSize: theme.fontSizes.md }}
          loading={createOrderForm.isSubmitting}
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
