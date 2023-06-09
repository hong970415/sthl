import { Grid, Group, Pagination, Table, Text } from '@mantine/core'
import { IImgInfo } from '@/entities/imageinfo'
import { IPagingQuery } from '@/services'
import useTranslationData from '@/hooks/useTranslation'
import { showDate } from '@/utils/date'
import { showBytes } from '@/utils/byteFormat'
import SelectLimit from '../SelectLimit/SelectLimit'
import AlbumImg from './AlbumImg'
import useStyles from './AlbumList.style'

const fkeys = ['img', 'img_name', 'img_size', 'createdAt', 'updatedAt']
interface IAlbumListGalleryModeProps {
  data: IImgInfo[]
  show: boolean
  viewCol: 1 | 2 | 3 | 4 | 5 | 6
  onClick: (value: IImgInfo) => void
}
function AlbumListGalleryMode(props: IAlbumListGalleryModeProps) {
  const { data, show, viewCol, onClick } = props
  const { classes } = useStyles()

  return (
    <Grid py="md" columns={60} display={!show ? 'none' : undefined}>
      {data &&
        data.map((item) => {
          return (
            <Grid.Col
              key={item.id}
              span={Math.round(60 / viewCol)}
              className={classes.col}
              onClick={(event) => {
                onClick(item)
              }}
            >
              <AlbumImg imgSrc={item.imgUrl} />
            </Grid.Col>
          )
        })}
    </Grid>
  )
}
interface IAlbumListDetailModeProps {
  data: IImgInfo[]
  show: boolean
  onClick: (value: IImgInfo) => void
}
function AlbumListDetailMode(props: IAlbumListDetailModeProps) {
  const { data, show, onClick } = props
  const { t } = useTranslationData()
  const { classes } = useStyles()

  const rows = data.map((row) => (
    <tr
      key={row.id}
      onClick={() => {
        onClick(row)
      }}
      style={{ cursor: 'pointer' }}
    >
      <td style={{ width: '80px' }}>
        <AlbumImg imgSrc={row.imgUrl} />
      </td>
      <td>{row.imgName}</td>
      <td>{showBytes(parseInt(row.imgSize))}</td>
      <td>{showDate(row.createdAt)}</td>
      <td>{showDate(row.updatedAt)}</td>
    </tr>
  ))
  // console.log('tt:', t('album.img_size'))
  return (
    <Table striped highlightOnHover display={!show ? 'none' : undefined}>
      <thead>
        <tr>
          {fkeys.map((item) => {
            return <th key={item}>{t(`album.${item}`)}</th>
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
  )
}
interface IAlbumListProps {
  data: IImgInfo[]
  onClickImg: (value: IImgInfo) => void
  viewMode: string
  viewCol: 1 | 2 | 3 | 4 | 5 | 6
  filter: IPagingQuery
  total: number
  totalPage: number
  setQuery: (value: string) => void
  setPage: (value: number) => void
  setLimit: (value: number) => void
}
export default function AlbumList(props: IAlbumListProps) {
  const {
    data,
    onClickImg,
    viewCol = 4,
    viewMode,
    filter,
    total,
    totalPage,
    setQuery,
    setPage,
    setLimit,
  } = props
  const { t } = useTranslationData()
  const { classes } = useStyles()

  const handleLimitChange = (value: string | null) => {
    if (value) {
      setLimit(parseInt(value))
    }
  }
  const paginationEl = (
    <Group position="right" py="md">
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

  return (
    <>
      {paginationEl}
      <AlbumListGalleryMode
        data={data}
        show={viewMode === 'gallery'}
        viewCol={viewCol}
        onClick={onClickImg}
      />
      <AlbumListDetailMode
        data={data}
        show={viewMode === 'detail'}
        onClick={onClickImg}
      />
      {paginationEl}
    </>
  )
}
