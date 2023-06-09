import { Badge } from '@mantine/core'
import { GetOrderStatusLabelKey, OrderStatus } from '@/constants/constants'
import useTranslationData from '@/hooks/useTranslation'

interface IOrderStatusBadgeProps {
  value: string
}

function getOrderStatusColor(value: string) {
  switch (value) {
    case OrderStatus.Initiated:
      return 'gray'
    case OrderStatus.Confirmed:
      return 'blue'
    case OrderStatus.Shipping:
      return 'yellow'
    case OrderStatus.Canceled:
      return 'red'
    case OrderStatus.Completed:
      return 'green'
    default:
      return 'blue'
  }
}
export default function OrderStatusBadge(props: IOrderStatusBadgeProps) {
  const { value } = props
  const { t } = useTranslationData()
  const orderStatusStyle = {
    radius: 'xs',
    variant: 'outline',
    color: getOrderStatusColor(value),
  }

  return (
    <Badge size="lg" {...orderStatusStyle}>
      {t(GetOrderStatusLabelKey(value))}
    </Badge>
  )
}
