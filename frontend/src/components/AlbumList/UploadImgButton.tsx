import useTranslationData from '@/hooks/useTranslation'
import { API, StatusCode } from '@/services'
import { Button, FileButton } from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import { useRef, useState } from 'react'

interface IUploadImgButtonProps {
  refetch: () => void
}
export default function UploadImgButton(props: IUploadImgButtonProps) {
  const { refetch } = props
  const { t } = useTranslationData()
  const [state, setState] = useState({ isLoading: false })
  const resetRef = useRef<() => void>(null)

  const handleOnChange = async (value: File | null) => {
    if (!value) {
      showNotification({ color: 'red', message: '500' })
      return
    }
    const MIN_FILE_SIZE = 1024 * 8 // 8MB
    if (value.size / 1024 > MIN_FILE_SIZE) {
      showNotification({ color: 'red', message: 'Larger than 8Mb' })
      return
    }

    setState((prev) => ({ ...prev, isLoading: true }))
    const payload = new FormData()
    payload.append('file', value)
    const res = await API.postUploadAlbumImg(payload)
    resetRef.current?.()
    if (res.status === StatusCode.Created) {
      showNotification({ color: 'green', message: res.data.msg })
      refetch()
      setState((prev) => ({ ...prev, isLoading: false }))
      return
    }
    const message = res.errorMsg
    showNotification({ color: 'red', message: message })
    setState((prev) => ({ ...prev, isLoading: false }))
  }
  return (
    <FileButton
      resetRef={resetRef}
      onChange={handleOnChange}
      accept="image/png,image/jpeg"
    >
      {(props) => (
        <Button {...props} loading={state.isLoading}>
          {t('general.upload_img')}
        </Button>
      )}
    </FileButton>
  )
}
