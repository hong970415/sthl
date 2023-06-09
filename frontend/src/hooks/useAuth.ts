import { useSelector } from 'react-redux'
import { ReduxState } from '@/redux/store'

export default function useAuth() {
  return useSelector((reduxtState: ReduxState) => reduxtState.auth)
}
