export const OrderTableFieldKeys = [
  'id',
  'discount',
  'totalAmount',
  'remark',
  'status',
  'paymentStatus',
  'paymentMethod',
  'deliveryStatus',
  'shippingAddress',
  'trackingNumber',
  'createdAt',
  'updatedAt',
]
export const OrderDetailFieldKeys = [
  'id',
  'items',
  'discount',
  'totalAmount',
  'remark',
  'status',
  'paymentStatus',
  'paymentMethod',
  'deliveryStatus',
  'shippingAddress',
  'trackingNumber',
  'createdAt',
  'updatedAt',
]
export interface IOrderItem {
  id: string
  orderId: string
  productId: string
  purchasedName: string
  purchasedPrice: number
  quantity: number
}

export interface IOrder {
  id: string
  userId: string
  items: IOrderItem[]
  discount: number
  totalAmount: number
  remark: string
  status: string
  paymentStatus: string
  paymentMethod: string
  deliveryStatus: string
  shippingAddress: string
  trackingNumber: string
  isArchived: boolean
  createdAt: string
  updatedAt: string
}
