import { useFormik } from 'formik'
import { showNotification } from '@mantine/notifications'
import { Button, Flex, PasswordInput, useMantineTheme } from '@mantine/core'
import * as yup from 'yup'
import { API, StatusCode, IUserUpdatePwForm } from '@/services'
import useTranslationData from '@/hooks/useTranslation'

export interface IUserUpdatePwFormSchema extends IUserUpdatePwForm {
  confirmNewPassword: string
}

export const UserUpdatePwFormSchema = yup.object().shape({
  currentPassword: yup
    .string()
    .required('required_error')
    .min(8, 'signup_form_pw_length_error')
    .max(500, 'signup_form_pw_length_error'),
  newPassword: yup
    .string()
    .required('required_error')
    .min(8, 'signup_form_pw_length_error')
    .max(500, 'signup_form_pw_length_error'),
  confirmNewPassword: yup
    .string()
    .required('required_error')
    .oneOf([yup.ref('newPassword'), ''], 'signup_form_pw_not_match_error'),
})

export default function UserUpdatePwForm() {
  const theme = useMantineTheme()
  const { t } = useTranslationData()
  const updatePwForm = useFormik<IUserUpdatePwFormSchema>({
    initialValues: {
      currentPassword: '',
      newPassword: '',
      confirmNewPassword: '',
    },
    validationSchema: UserUpdatePwFormSchema,
    onSubmit: async (values, { resetForm }) => {
      // console.log('UserUpdatePwForm values', values)
      const payload = {
        currentPassword: values.currentPassword,
        newPassword: values.newPassword,
      }
      // console.log('UserUpdatePwForm payload', payload)
      const response = await API.putUpdatePassword(payload)
      if (response.success && response.status === StatusCode.Ok) {
        const message = response.data.msg
        resetForm()
        showNotification({ color: 'green', message: message })
      } else {
        const message = response.errorMsg
        showNotification({ color: 'red', message: message })
      }
    },
  })

  // actions
  const handleOnClickClear = (event: React.MouseEvent<HTMLButtonElement>) => {
    if (updatePwForm.dirty) {
      updatePwForm.resetForm()
    }
  }
  // console.log('updatePwForm', updatePwForm)
  return (
    <form onSubmit={updatePwForm.handleSubmit}>
      <PasswordInput
        mb="md"
        name="currentPassword"
        label={t('general.current_password')}
        inputWrapperOrder={['label', 'input', 'description', 'error']}
        value={updatePwForm.values.currentPassword}
        onChange={updatePwForm.handleChange}
        error={
          updatePwForm.errors.currentPassword &&
          updatePwForm.touched.currentPassword
            ? t(`error.${updatePwForm.errors.currentPassword}`)
            : null
        }
        autoComplete="off"
      />
      <PasswordInput
        mb="md"
        name="newPassword"
        label={t('general.new_password')}
        inputWrapperOrder={['label', 'input', 'description', 'error']}
        value={updatePwForm.values.newPassword}
        onChange={updatePwForm.handleChange}
        error={
          updatePwForm.errors.newPassword && updatePwForm.touched.newPassword
            ? t(`error.${updatePwForm.errors.newPassword}`)
            : null
        }
        autoComplete="off"
      />
      <PasswordInput
        mb="md"
        name="confirmNewPassword"
        label={t('general.confirm_new_password')}
        inputWrapperOrder={['label', 'input', 'description', 'error']}
        value={updatePwForm.values.confirmNewPassword}
        onChange={updatePwForm.handleChange}
        error={
          updatePwForm.errors.confirmNewPassword &&
          updatePwForm.touched.confirmNewPassword
            ? t(`error.${updatePwForm.errors.confirmNewPassword}`)
            : null
        }
        autoComplete="off"
      />
      <Flex justify="center" align="center" gap={'lg'}>
        <Button
          mt="md"
          mb="md"
          type="submit"
          sx={{ fontSize: theme.fontSizes.md }}
          loading={updatePwForm.isSubmitting}
        >
          {t('general.update')}
        </Button>
        <Button
          //
          mt="md"
          mb="md"
          variant="outline"
          sx={{ fontSize: theme.fontSizes.md }}
          color="gray"
          onClick={handleOnClickClear}
          disabled={updatePwForm.isSubmitting}
        >
          {t('general.clear')}
        </Button>
      </Flex>
    </form>
  )
}
