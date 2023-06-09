import { useFormik } from 'formik'
import * as yup from 'yup'
import {
  Button,
  PasswordInput,
  TextInput,
  useMantineTheme,
} from '@mantine/core'
import { showNotification } from '@mantine/notifications'
import { API, StatusCode } from '@/services'
import useTranslationData from '@/hooks/useTranslation'

export const SignupFormSchema = yup.object().shape({
  email: yup
    .string()
    .required('required_error')
    .email('login_form_email_error')
    .min(6, 'signup_form_email_length_error')
    .max(255, 'signup_form_email_length_error'),
  password: yup
    .string()
    .required('required_error')
    .min(6, 'signup_form_pw_length_error')
    .max(255, 'signup_form_pw_length_error'),
  confirmPassword: yup
    .string()
    .required('required_error')
    .oneOf([yup.ref('password'), ''], 'signup_form_pw_not_match_error'),
})

export interface ISignupFormProps {
  onSuccess: () => void
}

export default function SignupForm(props: ISignupFormProps) {
  const { t } = useTranslationData()
  const theme = useMantineTheme()
  const signupForm = useFormik({
    initialValues: {
      email: '',
      password: '',
      confirmPassword: '',
    },
    validationSchema: SignupFormSchema,
    onSubmit: async (values) => {
      const payload = {
        email: values.email,
        password: values.password,
      }
      const response = await API.postSignup(payload)
      if (response.success && response.status === StatusCode.Created) {
        const message = response.data.msg
        showNotification({ color: 'green', message: message })
        props.onSuccess()
      } else {
        const message = response.errorMsg
        showNotification({ color: 'red', message: message })
      }
    },
  })

  return (
    <form onSubmit={signupForm.handleSubmit}>
      <TextInput
        mb="md"
        name="email"
        label={t('general.email')}
        value={signupForm.values.email}
        onChange={signupForm.handleChange}
        error={
          signupForm.errors.email && signupForm.touched.email
            ? t(`error.${signupForm.errors.email}`)
            : null
        }
      />
      <PasswordInput
        mb="md"
        name="password"
        label={t('general.password')}
        value={signupForm.values.password}
        onChange={signupForm.handleChange}
        error={
          signupForm.errors.password && signupForm.touched.password
            ? t(`error.${signupForm.errors.password}`)
            : null
        }
        autoComplete="off"
      />
      <PasswordInput
        mb="md"
        name="confirmPassword"
        label={t('general.confirm_password')}
        value={signupForm.values.confirmPassword}
        onChange={signupForm.handleChange}
        error={
          signupForm.errors.confirmPassword &&
          signupForm.touched.confirmPassword
            ? t(`error.${signupForm.errors.confirmPassword}`)
            : null
        }
        autoComplete="off"
      />
      <Button
        mt="lg"
        mb="md"
        type="submit"
        sx={{ fontSize: theme.fontSizes.md }}
        loading={signupForm.isSubmitting}
        fullWidth
      >
        {t('general.signup')}
      </Button>
    </form>
  )
}
