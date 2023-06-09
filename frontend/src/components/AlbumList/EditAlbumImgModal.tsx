import { MouseEvent, useRef, useState } from 'react'
import { showNotification } from '@mantine/notifications'
import { Button, FileButton, Group, Modal, TextInput } from '@mantine/core'
import { IImgInfo } from '@/entities/imageinfo'
import { showBytes } from '@/utils/byteFormat'
import AlbumImg from './AlbumImg'
import { API, StatusCode } from '@/services'

interface IEditAlbumImgModalProps {
  opened: boolean
  onClose: () => void
  data: IImgInfo | null
  refetch: () => void
}
interface IEditAlbumImgModalState {
  isEditing: boolean
  isUpdating: boolean
  formImgFile: File | null
  formImgFilePreviewUrl: string | null
}
export default function EditAlbumImgModal(props: IEditAlbumImgModalProps) {
  const { opened, onClose, data, refetch } = props
  const resetRef = useRef<() => void>(null)
  const [state, setState] = useState<IEditAlbumImgModalState>({
    isEditing: false,
    isUpdating: false,
    formImgFile: null,
    formImgFilePreviewUrl: null,
  })

  const handleOnClickChangeImg = (value: File | null) => {
    if (!value || !data) {
      showNotification({ color: 'red', message: '500' })
      return
    }
    const renamedFile = new File([value], data.imgName, { type: value.type })
    const MIN_FILE_SIZE = 1024 * 8 // 8MB
    if (renamedFile.size / 1024 > MIN_FILE_SIZE) {
      showNotification({ color: 'red', message: 'Larger than 8Mb' })
      return
    }

    let reader = new FileReader()
    reader.onloadend = (event) => {
      const result =
        event &&
        event.target &&
        event.target?.result &&
        event.target?.result.toString()
      if (result) {
        setState((prev) => ({
          ...prev,
          formImgFile: value,
          formImgFilePreviewUrl: result,
        }))
      }
    }
    reader.readAsDataURL(value)
  }
  const handleOnClickResetImg = (event: MouseEvent<HTMLButtonElement>) => {
    resetRef.current?.()
    setState((prev) => ({
      ...prev,
      formImgFile: null,
      formImgFilePreviewUrl: null,
    }))
  }
  const handleOnClickUpload = async (event: MouseEvent<HTMLButtonElement>) => {
    // console.log('handleOnClickUpload state:', state)
    if (!data || !state.formImgFile) {
      showNotification({ color: 'red', message: '500' })
      return
    }
    setState((prev) => ({ ...prev, isUpdating: true }))
    const payload = new FormData()
    payload.append('file', state.formImgFile)
    const res = await API.putUpdateImgDataById(data.id, payload)
    // console.log('handleOnClickUpload res', res)
    if (res.status === StatusCode.Ok) {
      showNotification({ color: 'green', message: res.data.msg })
      refetch()
      setState((prev) => ({ ...prev, isUpdating: false }))
      handleOnCloseModal()
      return
    }
    const message = res.errorMsg
    showNotification({ color: 'red', message: message })
    setState((prev) => ({ ...prev, isUpdating: false }))
  }
  const handleOnClickCancel = (event: MouseEvent<HTMLButtonElement>) => {
    handleOnClickResetImg(event)
    setState((prev) => ({
      ...prev,
      isEditing: false,
    }))
  }
  const handleOnClickEdit = (event: MouseEvent<HTMLButtonElement>) => {
    setState((prev) => ({ ...prev, isEditing: true }))
  }
  const handleOnCloseModal = () => {
    resetRef.current?.()
    setState((prev) => ({
      ...prev,
      formImgFile: null,
      formImgFilePreviewUrl: null,
      isEditing: false,
    }))
    onClose()
  }
  const changeImgButtonEl = (
    <Group position="right" pt={'md'}>
      <FileButton
        resetRef={resetRef}
        onChange={handleOnClickChangeImg}
        accept="image/png,image/jpeg"
      >
        {(props) => (
          <Button
            {...props}
            // loading={state.isLoading}
            disabled={!state.isEditing}
            compact
          >
            Change
          </Button>
        )}
      </FileButton>
      <Button
        compact
        disabled={!state.isEditing || !state.formImgFilePreviewUrl}
        variant="outline"
        onClick={handleOnClickResetImg}
      >
        Reset
      </Button>
    </Group>
  )
  const actionButtonEl = (
    <Group position="center" pt="md">
      {state.isEditing ? (
        <>
          <Button onClick={handleOnClickUpload} disabled={!state.formImgFile}>
            Update
          </Button>
          <Button onClick={handleOnClickCancel} variant="outline">
            Cancel
          </Button>
        </>
      ) : (
        <Button onClick={handleOnClickEdit}>Edit</Button>
      )}
    </Group>
  )

  return (
    <Modal opened={opened} onClose={handleOnCloseModal} title="">
      {data && (
        <>
          <AlbumImg
            imgSrc={
              state.formImgFilePreviewUrl
                ? state.formImgFilePreviewUrl
                : data.imgUrl
            }
            hoverEffect={false}
          />
          {changeImgButtonEl}
          <TextInput
            label={'Image name:'}
            value={data.imgName}
            disabled
            readOnly
          />
          <TextInput
            label={'Image size:'}
            value={showBytes(
              state.formImgFile
                ? state.formImgFile.size
                : parseInt(data.imgSize)
            )}
            pt="md"
            disabled
            readOnly
          />
          {actionButtonEl}
        </>
      )}
    </Modal>
  )
}
