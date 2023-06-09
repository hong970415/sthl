import { IImgInfo } from '@/entities/imageinfo'
import useTranslationData from '@/hooks/useTranslation'
import { Grid, Select, SelectItem, Text } from '@mantine/core'
import { forwardRef } from 'react'
import AlbumImg from '../AlbumList/AlbumImg'

type ItemProps = React.ComponentPropsWithoutRef<'div'> &
  Pick<IImgInfo, 'imgUrl'> & {
    label: string
  }

const SelectItem = forwardRef<HTMLDivElement, ItemProps>( // eslint-disable-line react/display-name
  ({ label, imgUrl, ...others }: ItemProps, ref) => {
    return (
      <div ref={ref} {...others}>
        <Grid>
          <Grid.Col span={1}>
            <AlbumImg imgSrc={imgUrl} hoverEffect={false} />
          </Grid.Col>
          <Grid.Col span={9}>
            <Text>{label}</Text>
          </Grid.Col>
        </Grid>
      </div>
    )
  }
)
interface ISelectImgProps {
  data: readonly (string | SelectItem)[]
  label?: string | null
  name?: string
  value?: string
  onChange?: (value: string | null) => void
  disabled?: boolean
}
export default function SelectImg(props: ISelectImgProps) {
  const { data, label, name, value, onChange, disabled } = props
  const handleOnChange = (value: string | null) => {
    onChange && onChange(value)
  }
  return (
    <Select
      label={label}
      name={name}
      itemComponent={SelectItem}
      data={data}
      maxDropdownHeight={400}
      value={value}
      onChange={handleOnChange}
      disabled={disabled}
    />
  )
}
