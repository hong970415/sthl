import { IOrderItem } from '@/entities/order'
import { ICreateOrderItem } from '@/services'

export function calculateOrderTotalAmount(
  items: Array<ICreateOrderItem | IOrderItem>,
  discount: number
): number {
  const subTotal = items.reduce(
    (accumulator, currentValue) =>
      accumulator + currentValue.purchasedPrice * (currentValue.quantity || 0),
    0
  )
  const total = subTotal * discount
  return total
}

export function getFormattedOrderTotalAmount(
  items: Array<ICreateOrderItem | IOrderItem>,
  discount: number
) {
  return calculateOrderTotalAmount(items, discount).toFixed(2)
}
