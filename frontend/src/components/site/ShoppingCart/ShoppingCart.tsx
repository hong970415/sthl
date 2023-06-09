import { MouseEvent, useState } from 'react'
import { useRouter } from 'next/router'
import { showNotification } from '@mantine/notifications'
import {
  ActionIcon,
  Button,
  Divider,
  Group,
  Indicator,
  Popover,
  Text,
  Textarea,
} from '@mantine/core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faShoppingCart } from '@fortawesome/free-solid-svg-icons'
import { getPrice } from '@/utils/price'
import { calculateOrderTotalAmount } from '@/utils/order'
import { API, ICreateOrderForm } from '@/services'
import useTranslationData from '@/hooks/useTranslation'
import useUserSite from '@/hooks/userUserSite'
import ShoppingCartItem from './ShoppingCartItem'
import useStyles from './ShoppingCart.style'

interface IShoppingCartProps {}
export default function ShoppingCart(props: IShoppingCartProps) {
  const router = useRouter()
  const {} = props
  const { t } = useTranslationData()
  const { userSiteState } = useUserSite()
  const { classes } = useStyles({})
  const [state, setState] = useState({
    isPlacingOrder: false,
    shippingAddress: '',
  })

  const handleOnClickPlaceOrder = async (event: MouseEvent) => {
    if (userSiteState.editMode) {
      return
    }
    const payload: ICreateOrderForm = {
      remark: '',
      discount: 1,
      totalAmount: calculateOrderTotalAmount(
        userSiteState.cartItems.map((el) => ({
          productId: el.id,
          purchasedName: el.name,
          purchasedPrice: el.price,
          quantity: el.quantity,
        })),
        1
      ),
      paymentMethod: 'card',
      shippingAddress: state.shippingAddress,
      items: userSiteState.cartItems.map((el) => {
        return {
          productId: el.id,
          purchasedName: el.name,
          purchasedPrice: el.price,
          quantity: el.quantity,
        }
      }),
    }
    // console.log('payload', payload)
    setState((prev) => ({ ...prev, isPlacingOrder: true }))
    const res = await API.postCreateOrder(userSiteState.userId, payload)
    if (res.success) {
      showNotification({ color: 'green', message: res.data.msg })
      setTimeout(() => {
        router.reload()
      }, 400)
      return
    }
    const message = res.errorMsg
    showNotification({ color: 'red', message: message })
    setState((prev) => ({
      ...prev,
      isPlacingOrder: false,
    }))
  }

  const targetEl = (
    <Indicator
      size={16}
      color="red"
      label={
        userSiteState.cartItems.length > 9
          ? '9+'
          : userSiteState.cartItems.length
      }
      disabled={userSiteState.cartItems.length === 0}
    >
      <ActionIcon size={'lg'} variant={'light'} color="blue">
        <FontAwesomeIcon icon={faShoppingCart} />
      </ActionIcon>
    </Indicator>
  )
  const dropdownEl = (
    <div>
      <div className={classes.itemsContainer}>
        {userSiteState.cartItems.length === 0 ? (
          <Text c="dimmed" ta="center" size={'lg'}>
            Empty...
          </Text>
        ) : (
          userSiteState.cartItems.map((cartItem) => {
            return <ShoppingCartItem key={cartItem.id} cartItem={cartItem} />
          })
        )}
      </div>
      <Divider my="sm" />
      <Group position="right" pb="md">
        <Text>
          Total: $
          {getPrice(
            calculateOrderTotalAmount(
              userSiteState.cartItems.map((el) => ({
                productId: el.id,
                purchasedName: el.name,
                purchasedPrice: el.price,
                quantity: el.quantity,
              })),
              1
            )
          )}
        </Text>
      </Group>
      <Textarea
        mb={'md'}
        name="shippingAddress"
        label={t('order.shippingAddress')}
        value={state.shippingAddress}
        onChange={(event) => {
          setState((prev) => ({ ...prev, shippingAddress: event.target.value }))
        }}
        withAsterisk
      />
      <Group noWrap>
        <Button
          fullWidth
          disabled={
            userSiteState.cartItems.length === 0 ||
            state.shippingAddress.length === 0
          }
          onClick={handleOnClickPlaceOrder}
          loading={state.isPlacingOrder}
        >
          Place order
        </Button>
      </Group>
    </div>
  )
  return (
    <Popover
      width={500}
      position="bottom"
      radius="lg"
      shadow="md"
      withArrow
      withinPortal
    >
      <Popover.Target>{targetEl}</Popover.Target>
      <Popover.Dropdown p="md">{dropdownEl}</Popover.Dropdown>
    </Popover>
  )
}
