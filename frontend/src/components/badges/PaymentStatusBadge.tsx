import { Badge } from '@mantine/core'
import useTranslationData from '@/hooks/useTranslation'
import { GetPaymentStatusLabelKey, PaymentStatus } from '@/constants/constants'

interface IPaymentStatusBadgeProps {
  value: string
}
function getPaymentStatusColor(value: string) {
  switch (value) {
    case PaymentStatus.Pending:
      return 'yellow'
    case PaymentStatus.Fail:
      return 'red'
    case PaymentStatus.Paid:
      return 'green'
    case PaymentStatus.Refunded:
      return 'violet'
    case PaymentStatus.PartialRefunded:
      return 'violet'
    case PaymentStatus.NoRefund:
      return 'violet'
    case PaymentStatus.Voided:
      return 'gray'
    default:
      return 'gray'
  }
}
export default function PaymentStatusBadge(props: IPaymentStatusBadgeProps) {
  const { value } = props
  const { t } = useTranslationData()
  const paymentStatusStyle = {
    radius: 'xl',
    variant: 'filled',
    color: getPaymentStatusColor(value),
  }
  return (
    <Badge size="lg" {...paymentStatusStyle}>
      {t(GetPaymentStatusLabelKey(value))}
    </Badge>
  )
}
