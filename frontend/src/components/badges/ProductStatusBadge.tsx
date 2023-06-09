import { Badge } from '@mantine/core'
import { GetProductStatusLabelKey, ProductStatus } from '@/constants/constants'
import useTranslationData from '@/hooks/useTranslation'

interface IProductStatusBadgeProps {
  value: string
}
function getProductStatusColor(value: string) {
  switch (value) {
    case ProductStatus.Initiated:
      return 'gray'
    case ProductStatus.Active:
      return 'blue'
    case ProductStatus.OutOfStock:
      return 'yellow'
    case ProductStatus.Inactive:
      return 'red'
    default:
      return 'gray'
  }
}
export default function ProductStatusBadge(props: IProductStatusBadgeProps) {
  const { value } = props
  const { t } = useTranslationData()
  const productStatusStyle = {
    radius: 'xl',
    variant: 'outline',
    color: getProductStatusColor(value),
  }
  return (
    <Badge size="lg" {...productStatusStyle}>
      {t(GetProductStatusLabelKey(value))}
    </Badge>
  )
}
