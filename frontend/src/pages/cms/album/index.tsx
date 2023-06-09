import { GetStaticProps } from 'next'
import { useState } from 'react'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { faGrip, faList } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { Group, SegmentedControl, Select } from '@mantine/core'
import { IImgInfo } from '@/entities/imageinfo'
import useGetAlbumImgsByUserId from '@/hooks/api/useGetAlbumImgsByUserId'
import PageContentContainer from '@/components/PageContentContainer/PageContentContainer'
import AlbumList from '@/components/AlbumList/AlbumList'
import UploadImgButton from '@/components/AlbumList/UploadImgButton'
import EditAlbumImgModal from '@/components/AlbumList/EditAlbumImgModal'

/** getStaticProps (NextJS)
 * only on server-side
 * @param locale the type of locale
 * @returns
 */
export const getStaticProps: GetStaticProps = async (context) => {
  const { locale } = context!
  return {
    props: {
      ...(await serverSideTranslations(locale as string, ['data'])),
      // will be passed to the page component as props
    },
  }
}
interface IAlbumPageState {
  viewCol: 1 | 2 | 3 | 4 | 5 | 6
  viewMode: string
  isViewDetailModalOpen: boolean
  selectedImgInfo: IImgInfo | null
}

export default function GalleryPage() {
  const {
    filter,
    setQuery,
    setPage,
    setLimit,
    album,
    refetch,
    total,
    totalPage,
  } = useGetAlbumImgsByUserId({})
  const [state, setState] = useState<IAlbumPageState>({
    viewCol: 4,
    viewMode: 'detail',
    isViewDetailModalOpen: false,
    selectedImgInfo: null,
  })

  const handleOnClickImgData = (value: IImgInfo) => {
    setState((prev) => ({
      ...prev,
      isViewDetailModalOpen: true,
      selectedImgInfo: value,
    }))
  }
  const handleOnCloseModal = () => {
    setState((prev) => ({
      ...prev,
      isViewDetailModalOpen: false,
      selectedImgInfo: null,
    }))
  }

  const selectColmnsEl = (
    <Select
      width={'10px'}
      label={'Columns:'}
      data={['1', '2', '3', '4', '5', '6'].map((item) => ({
        label: item,
        value: item,
      }))}
      value={state.viewCol.toString()}
      onChange={(value) => {
        setState((prev) => ({
          ...prev,
          viewCol: parseInt(value!) as 1 | 2 | 3 | 4 | 5 | 6,
        }))
      }}
      size="sm"
    />
  )
  const selectModeEl = (
    <Group position="apart" align="end">
      <UploadImgButton refetch={refetch} />
      <Group position="right" align="end">
        {state.viewMode === 'gallery' ? selectColmnsEl : null}
        <SegmentedControl
          data={[
            { label: <FontAwesomeIcon icon={faGrip} />, value: 'gallery' },
            { label: <FontAwesomeIcon icon={faList} />, value: 'detail' },
          ]}
          value={state.viewMode}
          onChange={(value) => {
            setState((prev) => ({ ...prev, viewMode: value }))
          }}
        />
      </Group>
    </Group>
  )

  // console.log('state', state)
  return (
    <PageContentContainer>
      {selectModeEl}
      <EditAlbumImgModal
        opened={state.isViewDetailModalOpen}
        onClose={handleOnCloseModal}
        data={state.selectedImgInfo}
        refetch={refetch}
      />
      {album && album?.imgs && (
        <AlbumList
          data={album.imgs}
          viewMode={state.viewMode}
          viewCol={state.viewCol}
          onClickImg={handleOnClickImgData}
          // paging
          filter={filter}
          total={total}
          totalPage={totalPage}
          setQuery={setQuery}
          setPage={setPage}
          setLimit={setLimit}
        />
      )}
    </PageContentContainer>
  )
}
