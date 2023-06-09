import { useRouter } from 'next/router'
import { Select } from '@mantine/core'
import { faGlobe } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { GetLocalesWithLabel } from '@/utils/locale'

export default function SelectLocale() {
  const router = useRouter()
  const { locale, locales, pathname, asPath, query } = router

  const list = locales ? locales : []
  const handleChangeLocale = (value: string) => {
    router.push({ pathname, query }, asPath, { locale: value })
  }

  return (
    <Select
      data={GetLocalesWithLabel(list)}
      value={locale}
      onChange={handleChangeLocale}
      sx={{ width: '120px' }}
      icon={<FontAwesomeIcon icon={faGlobe} />}
    />
  )
}
