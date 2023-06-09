import { ReactNode } from 'react'
import { Paper } from '@mantine/core'

interface IPageContentContainerProps {
  children?: ReactNode
}
export default function PageContentContainer(
  props: IPageContentContainerProps
) {
  const { children } = props
  return (
    <Paper
      radius="md"
      withBorder
      p="lg"
      sx={(theme) => ({
        backgroundColor:
          theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.white,
      })}
    >
      {children}
    </Paper>
  )
}
