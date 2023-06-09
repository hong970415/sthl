import { MouseEvent } from 'react'
import { faMinus, faPlus } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { ActionIcon, Flex, Grid, Text } from '@mantine/core'
import { IProduct } from '@/entities/product'
import { getPrice } from '@/utils/price'
import useUserSite from '@/hooks/userUserSite'
import {
  createDecrementCartItemQuantityByIdAction,
  createIncrementCartItemQuantityByIdAction,
  createRemoveCartItemByIdAction,
} from '@/providers/UserSiteProvider/actions'
import useStyles from './ShoppingCartItem.style'

export default function ShoppingCartItem(props: { cartItem: IProduct }) {
  const { cartItem } = props
  const { userSiteStateDispatch } = useUserSite()
  const { classes } = useStyles({})

  const handleOnClickMinus = (event: MouseEvent) => {
    userSiteStateDispatch(
      createDecrementCartItemQuantityByIdAction(cartItem.id)
    )
  }
  const handleOnClickPlus = (event: MouseEvent) => {
    userSiteStateDispatch(
      createIncrementCartItemQuantityByIdAction(cartItem.id)
    )
  }
  const handleOnClickRemoveIcon = (event: MouseEvent) => {
    userSiteStateDispatch(createRemoveCartItemByIdAction(cartItem.id))
  }
  return (
    <div className={classes.root}>
      <Grid mb="sm">
        <Grid.Col span={3}>
          <img src={cartItem.imgUrl} className={classes.img} />
        </Grid.Col>

        <Grid.Col span={4}>
          <Flex direction={'column'}>
            <Text size={'lg'}>{cartItem.name}</Text>
            <Text size={'xs'} c="dimmed" lineClamp={2}>
              {cartItem.description}
            </Text>
          </Flex>
        </Grid.Col>

        <Grid.Col span={2}>
          <Flex align={'center'}>
            <ActionIcon
              color="red"
              radius="xl"
              variant="outline"
              size={'xs'}
              onClick={handleOnClickMinus}
              disabled={cartItem.quantity === 1}
            >
              <FontAwesomeIcon icon={faMinus} size="xs" />
            </ActionIcon>
            <Text pl="xs" pr="xs">
              {cartItem.quantity}
            </Text>
            <ActionIcon
              color="green"
              radius="xl"
              variant="outline"
              size={'xs'}
              onClick={handleOnClickPlus}
            >
              <FontAwesomeIcon icon={faPlus} size="xs" />
            </ActionIcon>
          </Flex>
        </Grid.Col>
        <Grid.Col span={2}>
          <Text>${getPrice(cartItem.quantity * cartItem.price)}</Text>
        </Grid.Col>
        <Grid.Col span={1}>
          <ActionIcon
            color="red"
            radius="xl"
            variant="filled"
            size={'sm'}
            onClick={handleOnClickRemoveIcon}
          >
            <FontAwesomeIcon icon={faMinus} />
          </ActionIcon>
        </Grid.Col>
      </Grid>
    </div>
  )
}
