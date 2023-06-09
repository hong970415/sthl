export function GetLocalesWithLabel(list: string[]) {
  return list.map((value) => {
    switch (value) {
      case 'en':
        return { value: value, label: 'Eng' }
      case 'zh-HK':
        return { value: value, label: '繁體中文' }
      default:
        return { value: value, label: '' }
    }
  })
}
