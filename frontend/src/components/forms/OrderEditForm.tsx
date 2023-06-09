import {
  GetDeliveryStatusLabelKey,
  GetDeliveryStatusListWithLabel,
  GetOrderStatusLabelKey,
  GetOrderStatusListWithLabel,
  GetPaymentMethodLabelKey,
  GetPaymentMethodListWithLabel,
  GetPaymentStatusLabelKey,
  GetPaymentStatusListWithLabel,
} from '@/constants/constants'
import {
  IOrder,
  OrderDetailFieldKeys,
  OrderTableFieldKeys,
} from '@/entities/order'
import useGetOrderById from '@/hooks/api/useGetOrderById'
import useTranslationData from '@/hooks/useTranslation'
import { showDate } from '@/utils/date'
import {
  ActionIcon,
  Button,
  Flex,
  NumberInput,
  Select,
  Skeleton,
  Text,
  Textarea,
  useMantineTheme,
} from '@mantine/core'
import { useFormik } from 'formik'
import * as yup from 'yup'
import { get } from 'lodash'
import { useState } from 'react'
import DeliveryStatusBadge from '../badges/DeliveryStatusBadge'
import OrderStatusBadge from '../badges/OrderStatusBadge'
import PaymentStatusBadge from '../badges/PaymentStatusBadge'
import {
  CreateOrderFormSchema,
  getOrderItemInitialValue,
} from './OrderCreateForm'
import { API, ICreateOrderItem, StatusCode } from '@/services'
import { showNotification } from '@mantine/notifications'
import { calculateOrderTotalAmount } from '@/utils/order'
import useGetProducts from '@/hooks/api/useGetProducts'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faMinus } from '@fortawesome/free-solid-svg-icons'

const UpdateOrderFormSchema = CreateOrderFormSchema.concat(
  yup.object().shape({
    status: yup.string().required('required_error'),
    paymentStatus: yup.string().required('required_error'),
    deliveryStatus: yup.string().required('required_error'),
  })
)
function getEditableOrderField(order: IOrder) {
  return {
    items: order.items,
    discount: order.discount,
    totalAmount: order.totalAmount,
    remark: order.remark,
    status: order.status,
    paymentStatus: order.paymentStatus,
    paymentMethod: order.paymentMethod,
    deliveryStatus: order.deliveryStatus,
    shippingAddress: order.shippingAddress,
  }
}
interface IOrderEditFormProps {
  orderId: string
}
export default function OrderEditForm(props: IOrderEditFormProps) {
  const { orderId } = props
  const theme = useMantineTheme()
  const { t } = useTranslationData()
  const { products, refetch: refetchProducts } = useGetProducts({ limit: 1000 })
  const { order, isLoading, refetch: refetchOrder } = useGetOrderById(orderId)
  const editOrderForm = useFormik({
    initialValues: {
      items: [getOrderItemInitialValue()],
      remark: '',
      discount: 1,
      // totalAmount: 0,
      paymentMethod: '',
      shippingAddress: '',
      status: '',
      paymentStatus: '',
      deliveryStatus: '',
    },

    validationSchema: UpdateOrderFormSchema,
    onSubmit: async (values) => {
      // console.log('ProductEditForm values', values)
      const payload = {
        items: values.items.map((el) => ({
          productId: el.productId,
          quantity: el.quantity,
          purchasedPrice: el.purchasedPrice,
          purchasedName: el.purchasedName,
        })),
        remark: values.remark,
        discount: values.discount,
        totalAmount: calculateOrderTotalAmount(values.items, values.discount),
        paymentMethod: values.paymentMethod,
        shippingAddress: values.shippingAddress,
        status: values.status,
        paymentStatus: values.paymentStatus,
        deliveryStatus: values.deliveryStatus,
        trackingNumber: order?.trackingNumber,
      }
      const response = await API.putUpdateOrderById(orderId, payload)
      if (response.success && response.status === StatusCode.Ok) {
        refetchOrder()
        setState((prev) => ({ ...prev, isEditing: false }))
      } else {
        const message = response.errorMsg
        showNotification({ color: 'red', message: message })
      }
    },
  })
  const [state, setState] = useState({ isEditing: false })

  // actions
  const handleOnClickEdit = (event: React.MouseEvent<HTMLButtonElement>) => {
    if (order) {
      editOrderForm.resetForm({
        values: getEditableOrderField(order),
      })
      setState((prev) => ({ ...prev, isEditing: true }))
    }
  }

  const handleOnClickCancel = (event: React.MouseEvent<HTMLButtonElement>) => {
    if (order) {
      setState((prev) => ({ ...prev, isEditing: false }))
      editOrderForm.resetForm({
        values: getEditableOrderField(order),
      })
    }
  }

  const handleOnChangeItem = (productId: string, index: number) => {
    const selectedProduct = findProductById(productId)
    const defaultValue = {
      productId: productId,
      quantity: 1,
      purchasedPrice: selectedProduct?.price,
      purchasedName: selectedProduct?.name,
    }
    editOrderForm.setFieldValue(`items[${index}]`, defaultValue)
  }
  const handleOnChangeItemQuantity = (index: number, value: number) => {
    editOrderForm.setFieldValue(`items[${index}].quantity`, value)
  }

  const handleOnChangeItemPrice = (index: number, value: number) => {
    editOrderForm.setFieldValue(`items[${index}].purchasedPrice`, value)
  }
  const handleOnChangePaymentMethod = (value: string) => {
    editOrderForm.setFieldValue('paymentMethod', value)
  }
  const handleOnChangePaymentStatus = (value: string) => {
    editOrderForm.setFieldValue('paymentStatus', value)
  }
  const handleOnChangeOrderStatus = (value: string) => {
    editOrderForm.setFieldValue('status', value)
  }
  const handleOnChangeDeliveryStatus = (value: string) => {
    editOrderForm.setFieldValue('deliveryStatus', value)
  }
  const handleOnChangeDiscount = (value: number) => {
    editOrderForm.setFieldValue('discount', value)
  }
  const handleRemoveItem = (productId: string) => {
    editOrderForm.setFieldValue(
      'items',
      editOrderForm.values.items.filter((item) => item.productId !== productId)
    )
  }
  const findProductById = (productId: string) => {
    return products.find((item) => item.id === productId)
  }

  const viewModeEl = (
    <Flex direction={'column'}>
      {OrderDetailFieldKeys.map((key) => {
        const fieldStyles = {
          size: 'lg',
          weight: 'bold',
          sx: { width: '20%' },
        }
        const contentStyles = {
          sx: { width: '80%' },
        }
        if (key === 'createdAt' || key === 'updatedAt') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`product.${key}`)}:</Text>
                <Text {...contentStyles}>{showDate(get(order, key, ''))}</Text>
              </Flex>
            </Skeleton>
          )
        } else if (key === 'items') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align="center" mb="md">
                <Text size="lg" weight="bold" sx={{ width: '20%' }}>
                  {t(`order.${key}`)}:
                </Text>
                <Flex align="center" sx={{ width: '80%' }}>
                  <Text sx={{ width: '100%' }} weight="bold">
                    {t('product.name')}
                  </Text>
                  <Text sx={{ width: '100%' }} weight="bold">
                    {t('product.quantity')}
                  </Text>
                  <Text sx={{ width: '100%' }} weight="bold">
                    {t('product.price')}
                  </Text>
                </Flex>
              </Flex>
              {order?.items &&
                order?.items.map((item) => {
                  return (
                    <Flex
                      key={item.productId}
                      mb="md"
                      align="center"
                      justify="end"
                      sx={{ width: '80%' }}
                      ml="auto"
                    >
                      <Text sx={{ width: '100%' }}>
                        {get(item, 'purchasedName', '')}
                      </Text>
                      <Text sx={{ width: '100%' }}>
                        {get(item, 'quantity', '')}
                      </Text>
                      <Text sx={{ width: '100%' }}>
                        {get(item, 'purchasedPrice', '')}
                      </Text>
                    </Flex>
                  )
                })}
            </Skeleton>
          )
        } else if (key === 'status') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Text {...contentStyles}>
                  <OrderStatusBadge value={get(order, key, '')} />
                </Text>
              </Flex>
            </Skeleton>
          )
        } else if (key === 'paymentStatus') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Text {...contentStyles}>
                  <PaymentStatusBadge value={get(order, key, '')} />
                </Text>
              </Flex>
            </Skeleton>
          )
        } else if (key === 'paymentMethod') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Text {...contentStyles}>
                  {t(GetPaymentMethodLabelKey(get(order, key, '')))}
                </Text>
              </Flex>
            </Skeleton>
          )
        } else if (key === 'deliveryStatus') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Text {...contentStyles}>
                  <DeliveryStatusBadge value={get(order, key, '')} />
                </Text>
              </Flex>
            </Skeleton>
          )
        }
        return (
          <Skeleton key={key} visible={isLoading}>
            <Flex align={'center'} mb="md">
              <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
              <Text {...contentStyles}>{get(order, key, '')}</Text>
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

  // get available product quantity after chosen product
  const getAvailableProductQuantity = (productItem: ICreateOrderItem) => {
    const flattedOrderItemByIdExcludeSelf = editOrderForm.values.items
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
  const formItemsEl =
    editOrderForm.values.items.length > 0
      ? editOrderForm.values.items.map((productItem, index) => {
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
                    editOrderForm,
                    `errors.items[${index}].productId`,
                    null
                  ) &&
                  get(editOrderForm, `touched.items[${index}].productId`, null)
                    ? t(
                        `error.${get(
                          editOrderForm,
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
                  availableProductQuantity
                    ? `${t('order.avail')} ${availableProductQuantity}`
                    : undefined
                }
                type="number"
                min={1}
                max={maxProductQuantity}
                value={productItem.quantity}
                onChange={(value) => {
                  handleOnChangeItemQuantity(index, value as number)
                }}
                error={
                  get(editOrderForm, `errors.items[${index}].quantity`, null) &&
                  get(editOrderForm, `touched.items[${index}].quantity`, null)
                    ? t(
                        `error.${get(
                          editOrderForm,
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
                    editOrderForm,
                    `errors.items[${index}].purchasedPrice`,
                    null
                  ) &&
                  get(
                    editOrderForm,
                    `touched.items[${index}].purchasedPrice`,
                    null
                  )
                    ? t(
                        `error.${get(
                          editOrderForm,
                          `errors.items[${index}].purchasedPrice`,
                          ''
                        )}`
                      )
                    : null
                }
                // disabled={!hasAdminPermission}
                sx={{ width: '100%' }}
                withAsterisk
              />
              {editOrderForm.values.items.length > 1
                ? removeItemButton(productItem.productId, index)
                : null}
            </Flex>
          )
        })
      : null
  const editModeEl = (
    <form onSubmit={editOrderForm.handleSubmit}>
      {OrderDetailFieldKeys.map((key) => {
        const fieldStyles = {
          size: 'lg',
          weight: 'bold',
          sx: { width: '20%' },
        }
        const contentStyles = {
          sx: { width: '80%' },
        }
        if (key === 'createdAt' || key === 'updatedAt') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`product.${key}`)}:</Text>
                <Text {...contentStyles}>{showDate(get(order, key, ''))}</Text>
              </Flex>
            </Skeleton>
          )
        } else if (key === 'items') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                {formItemsEl}
              </Flex>
            </Skeleton>
          )
        } else if (key === 'discount') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <NumberInput
                  name={'discount'}
                  min={0.4}
                  max={1}
                  precision={2}
                  step={0.01}
                  value={editOrderForm.values.discount}
                  onChange={handleOnChangeDiscount}
                  error={
                    editOrderForm.errors.discount &&
                    editOrderForm.touched.discount
                      ? t(`error.${editOrderForm.errors.discount}`)
                      : null
                  }
                  withAsterisk
                />
              </Flex>
            </Skeleton>
          )
        } else if (key === 'remark') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Textarea
                  name="remark"
                  value={editOrderForm.values.remark}
                  onChange={editOrderForm.handleChange}
                  error={
                    editOrderForm.errors.remark && editOrderForm.touched.remark
                      ? t(`error.${editOrderForm.errors.remark}`)
                      : null
                  }
                  withAsterisk
                  {...contentStyles}
                />
              </Flex>
            </Skeleton>
          )
        } else if (key === 'status') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Select
                  name={'status'}
                  data={GetOrderStatusListWithLabel(t)}
                  value={editOrderForm.values.status}
                  onChange={handleOnChangeOrderStatus}
                  error={
                    editOrderForm.errors.status && editOrderForm.touched.status
                      ? t(`error.${editOrderForm.errors.status}`)
                      : null
                  }
                  withAsterisk
                  withinPortal
                />
              </Flex>
            </Skeleton>
          )
        } else if (key === 'paymentStatus') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Select
                  name={'paymentStatus'}
                  data={GetPaymentStatusListWithLabel(t)}
                  value={editOrderForm.values.paymentStatus}
                  onChange={handleOnChangePaymentStatus}
                  error={
                    editOrderForm.errors.paymentStatus &&
                    editOrderForm.touched.paymentStatus
                      ? t(`error.${editOrderForm.errors.paymentStatus}`)
                      : null
                  }
                  withAsterisk
                  withinPortal
                />
              </Flex>
            </Skeleton>
          )
        } else if (key === 'paymentMethod') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Select
                  name={'paymentMethod'}
                  data={GetPaymentMethodListWithLabel(t)}
                  value={editOrderForm.values.paymentMethod}
                  onChange={handleOnChangePaymentMethod}
                  error={
                    editOrderForm.errors.paymentMethod &&
                    editOrderForm.touched.paymentMethod
                      ? t(`error.${editOrderForm.errors.paymentMethod}`)
                      : null
                  }
                  withAsterisk
                  withinPortal
                />
              </Flex>
            </Skeleton>
          )
        } else if (key === 'deliveryStatus') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Select
                  name={'deliveryStatus'}
                  data={GetDeliveryStatusListWithLabel(t)}
                  value={editOrderForm.values.deliveryStatus}
                  onChange={handleOnChangeDeliveryStatus}
                  error={
                    editOrderForm.errors.deliveryStatus &&
                    editOrderForm.touched.deliveryStatus
                      ? t(`error.${editOrderForm.errors.deliveryStatus}`)
                      : null
                  }
                  withAsterisk
                  withinPortal
                />
              </Flex>
            </Skeleton>
          )
        } else if (key === 'shippingAddress') {
          return (
            <Skeleton key={key} visible={isLoading}>
              <Flex align={'center'} mb="md">
                <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
                <Textarea
                  name="shippingAddress"
                  value={editOrderForm.values.shippingAddress}
                  onChange={editOrderForm.handleChange}
                  error={
                    editOrderForm.errors.shippingAddress &&
                    editOrderForm.touched.shippingAddress
                      ? t(`error.${editOrderForm.errors.shippingAddress}`)
                      : null
                  }
                  withAsterisk
                  {...contentStyles}
                />
              </Flex>
            </Skeleton>
          )
        }
        return (
          <Skeleton key={key} visible={isLoading}>
            <Flex align={'center'} mb="md">
              <Text {...fieldStyles}>{t(`order.${key}`)}:</Text>
              <Text {...contentStyles}>{get(order, key, '')}</Text>
            </Flex>
          </Skeleton>
        )
      })}
      <Flex justify="center" align="center" gap={'lg'}>
        <Button
          mt="md"
          mb="md"
          type="submit"
          sx={{ fontSize: theme.fontSizes.md }}
          loading={editOrderForm.isSubmitting}
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
          disabled={editOrderForm.isSubmitting}
        >
          {t('general.cancel')}
        </Button>
      </Flex>
    </form>
  )
  return state.isEditing ? editModeEl : viewModeEl
}
