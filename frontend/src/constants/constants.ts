enum ProductStatus {
  Initiated = 'initiated',
  Active = 'active',
  OutOfStock = 'outOfStock',
  Inactive = 'inactive',
}
function GetProductStatusLabelKey(value: string) {
  return `productStatus.${value}`
}
function GetProductStatusListWithLabel(t: (key: string) => string) {
  return [
    { value: ProductStatus.Initiated, label: t('productStatus.initiated') },
    { value: ProductStatus.Active, label: t('productStatus.active') },
    { value: ProductStatus.OutOfStock, label: t('productStatus.outOfStock') },
    { value: ProductStatus.Inactive, label: t('productStatus.inactive') },
  ]
}

enum OrderStatus {
  Initiated = 'initiated',
  Confirmed = 'confirmed',
  Shipping = 'shipping',
  Canceled = 'canceled',
  Completed = 'completed',
}
function GetOrderStatusLabelKey(value: string) {
  return `orderStatus.${value}`
}
function GetOrderStatusListWithLabel(t: (key: string) => string) {
  return [
    { value: OrderStatus.Initiated, label: t('orderStatus.initiated') },
    { value: OrderStatus.Confirmed, label: t('orderStatus.confirmed') },
    { value: OrderStatus.Shipping, label: t('orderStatus.shipping') },
    { value: OrderStatus.Canceled, label: t('orderStatus.canceled') },
    { value: OrderStatus.Completed, label: t('orderStatus.completed') },
  ]
}

enum PaymentStatus {
  Pending = 'pending',
  Fail = 'fail',
  Paid = 'paid',
  Refunded = 'refunded',
  PartialRefunded = 'partialRefunded',
  NoRefund = 'noRefund',
  Voided = 'voided',
}
function GetPaymentStatusLabelKey(value: string) {
  return `paymentStatus.${value}`
}
function GetPaymentStatusListWithLabel(t: (key: string) => string) {
  return [
    { value: PaymentStatus.Pending, label: t('paymentStatus.pending') },
    { value: PaymentStatus.Fail, label: t('paymentStatus.fail') },
    { value: PaymentStatus.Paid, label: t('paymentStatus.paid') },
    { value: PaymentStatus.Refunded, label: t('paymentStatus.refunded') },
    {
      value: PaymentStatus.PartialRefunded,
      label: t('paymentStatus.partialRefunded'),
    },
    { value: PaymentStatus.NoRefund, label: t('paymentStatus.noRefund') },
    { value: PaymentStatus.Voided, label: t('paymentStatus.voided') },
  ]
}

enum PaymentMethod {
  Card = 'card',
}
function GetPaymentMethodLabelKey(value: string) {
  return `paymentMethod.${value}`
}
function GetPaymentMethodListWithLabel(t: (key: string) => string) {
  return [{ value: PaymentMethod.Card, label: t('paymentMethod.card') }]
}

enum DeliveryStatus {
  Pending = 'pending',
  Shipping = 'shipping',
  Completed = 'completed',
}
function GetDeliveryStatusLabelKey(value: string) {
  return `deliveryStatus.${value}`
}
function GetDeliveryStatusListWithLabel(t: (key: string) => string) {
  return [
    { value: DeliveryStatus.Pending, label: t('deliveryStatus.pending') },
    { value: DeliveryStatus.Shipping, label: t('deliveryStatus.shipping') },
    { value: DeliveryStatus.Completed, label: t('deliveryStatus.completed') },
  ]
}

export {
  ProductStatus,
  GetProductStatusLabelKey,
  GetProductStatusListWithLabel,
  OrderStatus,
  GetOrderStatusLabelKey,
  GetOrderStatusListWithLabel,
  PaymentStatus,
  GetPaymentStatusLabelKey,
  GetPaymentStatusListWithLabel,
  PaymentMethod,
  GetPaymentMethodLabelKey,
  GetPaymentMethodListWithLabel,
  DeliveryStatus,
  GetDeliveryStatusLabelKey,
  GetDeliveryStatusListWithLabel,
}
