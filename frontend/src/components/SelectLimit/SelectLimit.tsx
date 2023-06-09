import { Select } from '@mantine/core'

export const LimitList = [
  { value: '10', label: '10' },
  { value: '20', label: '20' },
  { value: '50', label: '50' },
  { value: '100', label: '100' },
]

export interface ISelectLimitProps {
  value: string
  onChange: (value: string | null) => void
}

export default function SelectLimit(props: ISelectLimitProps) {
  const { value, onChange } = props

  const handleOnChange = (value: string | null) => {
    onChange && onChange(value)
  }
  return (
    <Select
      value={value}
      onChange={handleOnChange}
      data={LimitList}
      sx={{ width: 100 }}
    />
  )
}
