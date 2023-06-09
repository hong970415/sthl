import { ChangeEvent, MouseEvent, RefObject, useState } from 'react'
import { Button, ColorInput, Divider, Flex, TextInput } from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import {
  createUpdateHomepageImgUrlAction,
  createUpdateHomepageTextAction,
  createUpdateHomepageTextColorAction,
  createUpdateSitenameAction,
  UserSiteStateActionType,
} from '@/providers/UserSiteProvider/actions'
import useUserSite from '@/hooks/userUserSite'
import { API, StatusCode } from '@/services'
import useGetAlbumImgsByUserId from '@/hooks/api/useGetAlbumImgsByUserId'
import SelectImg from '@/components/SelectImg/SelectImg'
import useTranslationData from '@/hooks/useTranslation'

interface IDesignPanelProps {
  sitePath: string
  iframeRef: RefObject<HTMLIFrameElement>
  refetchSiteUi: () => void
  reloadIframe: () => void
}
export default function DesignPanel(props: IDesignPanelProps) {
  const { sitePath, iframeRef, refetchSiteUi, reloadIframe } = props
  const { userSiteState, userSiteStateDispatch, resetUiData } = useUserSite()
  const { album, isFetched } = useGetAlbumImgsByUserId({ limit: 1000 })
  const { t } = useTranslationData()
  const [state, setState] = useState({
    isEditing: false,
    isLoading: false,
  })

  const handleOnChangeSitename = (event: ChangeEvent<HTMLInputElement>) => {
    const value = event.currentTarget.value
    if (!value) {
      return
    }
    const action = createUpdateSitenameAction(value)
    userSiteStateDispatch(action)
    if (!iframeRef.current) return
    iframeRef.current.contentWindow?.postMessage(action)
  }
  const handleOnChangeHomepageImgUrl = (value: string | null) => {
    if (!value) {
      return
    }
    const action = createUpdateHomepageImgUrlAction(value)
    userSiteStateDispatch(action)
    if (!iframeRef.current) return
    iframeRef.current.contentWindow?.postMessage(action)
  }
  const handleOnChangeHomepageText = (event: ChangeEvent<HTMLInputElement>) => {
    const value = event.currentTarget.value
    const action = createUpdateHomepageTextAction(value)
    userSiteStateDispatch(action)
    if (!iframeRef.current) return
    iframeRef.current.contentWindow?.postMessage(action)
  }
  const handleOnChangeHomepageTextColor = (value: string) => {
    const action = createUpdateHomepageTextColorAction(value)
    userSiteStateDispatch(action)
    if (!iframeRef.current) return
    iframeRef.current.contentWindow?.postMessage(action)
  }

  const handleOnClickUpdate = async (event: MouseEvent<HTMLButtonElement>) => {
    setState((prev) => ({ ...prev, isLoading: true }))
    const payload = {
      sitename: userSiteState.uiData.sitename,
      homepageImgUrl: userSiteState.uiData.homepageImgUrl,
      homepageText: userSiteState.uiData.homepageText,
      homepageTextColor: userSiteState.uiData.homepageTextColor,
    }
    const res = await API.putUpsertUserSiteUiDataById(payload)
    // console.log('update res:', res)

    if (res.status === StatusCode.Ok) {
      const message = res.data.msg
      showNotification({ color: 'green', message: message })
      //
      refetchSiteUi()
      reloadIframe()
      setState((prev) => ({ ...prev, isEditing: false }))
    } else {
      const message = res.errorMsg
      showNotification({ color: 'red', message: message })
    }
    setState((prev) => ({ ...prev, isLoading: false }))
  }
  const handleOnClickEdit = (event: MouseEvent<HTMLButtonElement>) => {
    setState((prev) => ({ ...prev, isEditing: true }))
  }
  const handleOnClickCancel = (event: MouseEvent<HTMLButtonElement>) => {
    setState((prev) => ({ ...prev, isEditing: false }))
    resetUiData()
    if (!iframeRef.current) return
    iframeRef.current.contentWindow?.postMessage({
      type: UserSiteStateActionType.ResetUiData,
    })
  }
  const editformEl = (
    <>
      <TextInput
        label={t('site.sitename')}
        value={userSiteState.uiData.sitename}
        onChange={handleOnChangeSitename}
        disabled={!state.isEditing}
      />
      <SelectImg
        label={t('site.hero_img')}
        data={
          album && album.imgs
            ? album.imgs.map((el) => ({
                value: el.imgUrl,
                label: el.imgName,
                imgUrl: el.imgUrl,
              }))
            : []
        }
        onChange={handleOnChangeHomepageImgUrl}
        value={userSiteState.uiData.homepageImgUrl}
        disabled={!state.isEditing}
      />
      <TextInput
        label={t('site.hero_title')}
        placeholder="Welcome"
        value={userSiteState.uiData.homepageText}
        onChange={handleOnChangeHomepageText}
        disabled={!state.isEditing}
      />
      <ColorInput
        label={t('site.hero_title_color')}
        placeholder="Pick color"
        value={userSiteState.uiData.homepageTextColor}
        onChange={handleOnChangeHomepageTextColor}
        disabled={!state.isEditing}
      />
      <Flex direction={{ base: 'column', md: 'row' }} py="lg" gap="sm">
        {state.isEditing ? (
          <Button
            fullWidth
            onClick={handleOnClickUpdate}
            loading={state.isLoading}
          >
            {t('general.update')}
          </Button>
        ) : (
          <Button fullWidth onClick={handleOnClickEdit}>
            {t('general.edit')}
          </Button>
        )}
        <Button
          fullWidth
          variant={'outline'}
          disabled={!state.isEditing}
          onClick={handleOnClickCancel}
        >
          {t('general.cancel')}
        </Button>
      </Flex>
    </>
  )

  return (
    <div>
      <a target="_blank" href={sitePath}>
        <Button fullWidth color="teal">
          {t('site.view_your_site')}
        </Button>
      </a>
      <Divider my="md" />
      {editformEl}
    </div>
  )
}
