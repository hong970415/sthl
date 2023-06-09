import { Badge } from '@mantine/core'
import {
  DeliveryStatus,
  GetDeliveryStatusLabelKey,
} from '@/constants/constants'
import useTranslationData from '@/hooks/useTranslation'

interface IDeliveryStatusBadgeProps {
  value: string
}
function getDeliveryStatusColor(value: string) {
  switch (value) {
    case DeliveryStatus.Pending:
      return 'yellow'
    case DeliveryStatus.Shipping:
      return 'blue'
    case DeliveryStatus.Completed:
      return 'green'
    default:
      return 'blue'
  }
}
export default function DeliveryStatusBadge(props: IDeliveryStatusBadgeProps) {
  const { value } = props
  const { t } = useTranslationData()
  const deliveryStatusStyle = {
    radius: 'xl',
    variant: 'dot',
    color: getDeliveryStatusColor(value),
  }
  return (
    <Badge size="lg" {...deliveryStatusStyle}>
      {t(GetDeliveryStatusLabelKey(value))}
    </Badge>
  )
}
