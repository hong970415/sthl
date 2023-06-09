import queryString from 'query-string'

export function makeQueryString(value: any) {
  return (
    '?' +
    queryString.stringify(value, { skipEmptyString: true, skipNull: true })
  )
}
