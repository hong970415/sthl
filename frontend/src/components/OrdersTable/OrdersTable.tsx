import { ChangeEvent, MouseEvent } from 'react'
import { useRouter } from 'next/router'
import {
  Button,
  Flex,
  Group,
  MediaQuery,
  Pagination,
  Paper,
  ScrollArea,
  Table,
  Text,
  TextInput,
  useMantineTheme,
} from '@mantine/core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faMagnifyingGlass, faPlus } from '@fortawesome/free-solid-svg-icons'
import { IPagingQuery } from '@/services'
import { IOrder, OrderTableFieldKeys } from '@/entities/order'
import useGetOrders from '@/hooks/api/useGetOrders'
import useTranslationData from '@/hooks/useTranslation'
import SelectLimit from '../SelectLimit/SelectLimit'
import { showDate } from '@/utils/date'
import { getPrice } from '@/utils/price'
import {
  GetDeliveryStatusLabelKey,
  GetOrderStatusLabelKey,
  GetPaymentMethodLabelKey,
  GetPaymentStatusLabelKey,
} from '@/constants/constants'
import OrderStatusBadge from '../badges/OrderStatusBadge'
import PaymentStatusBadge from '../badges/PaymentStatusBadge'
import DeliveryStatusBadge from '../badges/DeliveryStatusBadge'

interface IOrdersTableHeaderProps {
  pathToAddOrder: string
  filter: IPagingQuery
  setQuery: (value: string) => void
}
function OrdersTableHeader(props: IOrdersTableHeaderProps) {
  const { pathToAddOrder, filter, setQuery } = props
  const router = useRouter()
  const theme = useMantineTheme()
  const { t } = useTranslationData()

  const handleOnClickAdd = (event: MouseEvent<HTMLAnchorElement>) => {
    event.preventDefault()
    router.push(pathToAddOrder)
  }
  const handleOnChangeQuery = (event: ChangeEvent<HTMLInputElement>) => {
    const { value } = event.currentTarget
    setQuery(value)
  }
  const handleOnClickClear = (vent: MouseEvent<HTMLButtonElement>) => {
    setQuery('')
  }

  const addOrderButton = (
    <Button
      component="a"
      href={pathToAddOrder}
      variant="outline"
      sx={{ fontSize: theme.fontSizes.sm }}
      leftIcon={<FontAwesomeIcon icon={faPlus} />}
      onClick={handleOnClickAdd}
    >
      {t('general.create')}
    </Button>
  )
  return (
    <>
      <MediaQuery largerThan="sm" styles={{ display: 'none' }}>
        <Group position="right" mb="md">
          {addOrderButton}
        </Group>
      </MediaQuery>
      <Group position="apart">
        <MediaQuery smallerThan="sm" styles={{ width: '100%' }}>
          <Flex gap="sm" direction={{ base: 'column', sm: 'row' }}>
            <TextInput
              placeholder={t<string>('product.search_placeholder')}
              icon={<FontAwesomeIcon icon={faMagnifyingGlass} />}
              value={filter.query}
              onChange={handleOnChangeQuery}
            />
            <Button
              sx={{ fontSize: theme.fontSizes.sm }}
              variant="outline"
              color="gray"
              onClick={handleOnClickClear}
            >
              {t('general.clear')}
            </Button>
          </Flex>
        </MediaQuery>
        <MediaQuery smallerThan="sm" styles={{ display: 'none' }}>
          {addOrderButton}
        </MediaQuery>
      </Group>
    </>
  )
}

interface IOrdersTablePaginationProps {
  filter: IPagingQuery
  total: number
  totalPage: number
  setLimit: (value: number) => void
  setPage: (value: number) => void
}
function ProductsTablePagination(props: IOrdersTablePaginationProps) {
  const { filter, total, totalPage, setLimit, setPage } = props
  const { t } = useTranslationData()
  const handleLimitChange = (value: string | null) => {
    if (value) {
      setLimit(parseInt(value))
    }
  }
  return (
    <Group position="right">
      <Text align="center">
        {t('general.total')}:{total}
      </Text>
      <SelectLimit
        value={filter.limit.toString()}
        onChange={handleLimitChange}
      />
      <Pagination value={filter.page} total={totalPage} onChange={setPage} />
    </Group>
  )
}

interface IOrdersTableContentProps {
  orders: IOrder[]
  onClickRow?: (id: string) => void
}
function OrdersTableContent(props: IOrdersTableContentProps) {
  const { orders, onClickRow } = props
  const { t } = useTranslationData()

  const rows = orders.map((row) => (
    <tr
      key={row.id}
      onClick={() => {
        onClickRow && onClickRow(row.id)
      }}
      style={{ cursor: 'pointer' }}
    >
      <td>{row.id}</td>
      <td>{row.discount}</td>
      <td>${getPrice(row.totalAmount)}</td>
      <td>{row.remark}</td>
      <td>
        <OrderStatusBadge value={row.status} />
      </td>
      <td>
        <PaymentStatusBadge value={row.paymentStatus} />
      </td>
      <td>{t(GetPaymentMethodLabelKey(row.paymentMethod))}</td>
      <td>
        <DeliveryStatusBadge value={row.deliveryStatus} />
      </td>
      <td>{row.shippingAddress}</td>
      <td>{row.trackingNumber}</td>
      <td>{showDate(row.createdAt)}</td>
      <td>{showDate(row.updatedAt)}</td>
    </tr>
  ))
  const table = (
    <ScrollArea mb="md" mt="md" sx={{ height: '100%' }}>
      <Table striped highlightOnHover>
        <thead>
          <tr>
            {OrderTableFieldKeys.map((item) => {
              return <th key={item}>{t(`order.${item}`)}</th>
            })}
          </tr>
        </thead>
        <tbody>
          {rows.length > 0 ? (
            rows
          ) : (
            <tr>
              <td colSpan={12}>
                <Text weight={500} align="center">
                  {t('general.nothing_found')}
                </Text>
              </td>
            </tr>
          )}
        </tbody>
      </Table>
    </ScrollArea>
  )
  return table
}

export default function OrdersTable() {
  const router = useRouter()
  const { filter, setQuery, setPage, setLimit, orders, total, totalPage } =
    useGetOrders({})
  const pathToAddOrder = '/cms/order/add'
  const handleOnClickRow = (id: string) => {
    router.push(`/cms/order/detail?id=${id}`)
  }

  const tablePagination = (
    <ProductsTablePagination
      filter={filter}
      total={total}
      totalPage={totalPage}
      setLimit={setLimit}
      setPage={setPage}
    />
  )
  return (
    <>
      <OrdersTableHeader
        pathToAddOrder={pathToAddOrder}
        filter={filter}
        setQuery={setQuery}
      />
      <Paper
        withBorder
        radius="md"
        p="sm"
        mt="md"
        sx={(theme) => ({
          backgroundColor:
            theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.white,
        })}
      >
        {tablePagination}
        <OrdersTableContent orders={orders} onClickRow={handleOnClickRow} />
        {tablePagination}
      </Paper>
    </>
  )
}
