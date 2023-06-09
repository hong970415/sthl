import dayjs from 'dayjs'

export function showDate(value: string) {
  return dayjs(value).format('YYYY-MM-DD hh:mm:ss A')
}
