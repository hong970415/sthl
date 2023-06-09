import { useRouter } from 'next/router'
import { ChangeEvent, MouseEvent } from 'react'
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
import useGetProducts from '@/hooks/api/useGetProducts'
import useTranslationData from '@/hooks/useTranslation'
import { IProduct, ProductFieldKeys } from '@/entities/product'
import { showDate } from '@/utils/date'
import SelectLimit from '../SelectLimit/SelectLimit'
import { GetProductStatusLabelKey } from '@/constants/constants'
import ProductStatusBadge from '../badges/ProductStatusBadge'

interface IProductsTableHeaderProps {
  pathToAddProduct: string
  filter: IPagingQuery
  setQuery: (value: string) => void
}
function ProductsTableHeader(props: IProductsTableHeaderProps) {
  const { pathToAddProduct, filter, setQuery } = props
  const router = useRouter()
  const theme = useMantineTheme()
  const { t } = useTranslationData()

  const handleOnClickAdd = (event: MouseEvent<HTMLAnchorElement>) => {
    event.preventDefault()
    router.push(pathToAddProduct)
  }
  const handleOnChangeQuery = (event: ChangeEvent<HTMLInputElement>) => {
    const { value } = event.currentTarget
    setQuery(value)
  }
  const handleOnClickClear = (vent: MouseEvent<HTMLButtonElement>) => {
    setQuery('')
  }

  const addProductButton = (
    <Button
      component="a"
      href={pathToAddProduct}
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
          {addProductButton}
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
          {addProductButton}
        </MediaQuery>
      </Group>
    </>
  )
}

interface IProductsTablePaginationProps {
  filter: IPagingQuery
  total: number
  totalPage: number
  setLimit: (value: number) => void
  setPage: (value: number) => void
}
function ProductsTablePagination(props: IProductsTablePaginationProps) {
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

interface IProductsTableContentProps {
  products: IProduct[]
  onClickRow?: (id: string) => void
}
function ProductsTableContent(props: IProductsTableContentProps) {
  const { products, onClickRow } = props
  const { t } = useTranslationData()

  const rows = products.map((row) => (
    <tr
      key={row.id}
      onClick={() => {
        onClickRow && onClickRow(row.id)
      }}
      style={{ cursor: 'pointer' }}
    >
      <td>{row.id}</td>
      <td>{row.name}</td>
      <td>{row.price}</td>
      <td>{row.quantity}</td>
      <td>{row.description}</td>
      {/* <td>{t(GetProductStatusLabelKey(row.status))}</td> */}
      <td>
        <ProductStatusBadge value={row.status} />
      </td>
      <td>{showDate(row.createdAt)}</td>
      <td>{showDate(row.updatedAt)}</td>
    </tr>
  ))
  const table = (
    <ScrollArea mb="md" mt="md" sx={{ height: '100%' }}>
      <Table striped highlightOnHover>
        <thead>
          <tr>
            {ProductFieldKeys.map((item) => {
              return <th key={item}>{t(`product.${item}`)}</th>
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

export default function ProductsTable() {
  const router = useRouter()
  const { filter, setQuery, setPage, setLimit, products, total, totalPage } =
    useGetProducts({})
  const pathToAddProduct = '/cms/product/add'
  const handleOnClickRow = (id: string) => {
    router.push(`/cms/product/detail?id=${id}`)
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
      <ProductsTableHeader
        pathToAddProduct={pathToAddProduct}
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
        <ProductsTableContent
          products={products}
          onClickRow={handleOnClickRow}
        />
        {tablePagination}
      </Paper>
    </>
  )
}
