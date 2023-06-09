import { useContext } from 'react'
import { UserSiteContext } from '@/providers/UserSiteProvider/context'

export default function useUserSite() {
  const context = useContext(UserSiteContext)
  if (context === undefined) {
    throw new Error('useUserSite must be used within a UserSiteProvider')
  }
  return context!
}
