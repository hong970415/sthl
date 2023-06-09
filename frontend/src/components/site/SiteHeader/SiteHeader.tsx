import { MouseEvent, useState } from 'react'
import {
  Header,
  Container,
  Group,
  Burger,
  Paper,
  Transition,
} from '@mantine/core'
import { useDisclosure } from '@mantine/hooks'
import Link from 'next/link'
import useStyles from './SiteHeader.style'
import ShoppingCart from '../ShoppingCart/ShoppingCart'

interface SiteHeaderProps {
  height: string
  links: { link: string; label: string }[]
}
export default function SiteHeader(props: SiteHeaderProps) {
  const { height, links } = props
  const [opened, { toggle, close }] = useDisclosure(false)
  const [active, setActive] = useState(links[0].link)
  const { classes, cx } = useStyles({ headerHeight: height })

  const handleOnClickItem = (event: MouseEvent, value: string) => {
    setActive(value)
    close()
  }
  const items =
    links &&
    links.map((link) => (
      <Link
        key={link.label}
        href={`#${link.link}`}
        className={cx(classes.link, {
          [classes.linkActive]: active === link.link,
        })}
        onClick={(event) => {
          handleOnClickItem(event, link.link)
        }}
        scroll={false}
      >
        {link.label}
      </Link>
    ))

  const mobileVersionEl = (
    <>
      <Burger
        opened={opened}
        onClick={toggle}
        className={classes.burger}
        size="sm"
      />
      <Transition transition="pop-top-right" duration={200} mounted={opened}>
        {(styles) => (
          <Paper className={classes.dropdown} withBorder style={styles}>
            {items}
          </Paper>
        )}
      </Transition>
    </>
  )

  return (
    <Header height={height} className={classes.root}>
      <Container className={classes.header}>
        <Group spacing={5} className={classes.links}>
          {items}
        </Group>

        {mobileVersionEl}

        {/* carts */}
        <ShoppingCart />
      </Container>
    </Header>
  )
}
