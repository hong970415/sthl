import { useFormik } from 'formik'
import * as yup from 'yup'
import {
  Button,
  PasswordInput,
  TextInput,
  useMantineTheme,
} from '@mantine/core'
import { useAppDispatch } from '@/redux/store'
import { thunkLogin } from '@/redux/auth/thunk'
import useTranslationData from '@/hooks/useTranslation'

export const LoginFormSchema = yup.object().shape({
  email: yup
    .string()
    .required('required_error')
    .email('login_form_email_error')
    .min(6, 'login_form_email_length_error')
    .max(255, 'login_form_email_length_error'),
  password: yup
    .string()
    .required('required_error')
    .min(6, 'login_form_pw_length_error')
    .max(255, 'login_form_pw_length_error'),
})

export interface ILoginFormProps {}
export default function LoginForm(props: ILoginFormProps) {
  const { t } = useTranslationData()
  const theme = useMantineTheme()
  const dispatch = useAppDispatch()
  const loginForm = useFormik({
    initialValues: {
      email: '',
      password: '',
      // email: 'asd55@adawd.et',
      // password: 'qasdsadwx3',
    },
    validationSchema: LoginFormSchema,
    onSubmit: async (values) => {
      dispatch(thunkLogin(values))
    },
  })

  return (
    <form onSubmit={loginForm.handleSubmit}>
      <TextInput
        //
        mb="md"
        name="email"
        label={t('general.email')}
        inputWrapperOrder={['label', 'input', 'description', 'error']}
        value={loginForm.values.email}
        onChange={loginForm.handleChange}
        error={
          loginForm.errors.email && loginForm.touched.email
            ? t(`error.${loginForm.errors.email}`)
            : null
        }
      />
      <PasswordInput
        mb="md"
        name="password"
        label={t('general.password')}
        inputWrapperOrder={['label', 'input', 'description', 'error']}
        value={loginForm.values.password}
        onChange={loginForm.handleChange}
        error={
          loginForm.errors.password && loginForm.touched.password
            ? t(`error.${loginForm.errors.password}`)
            : null
        }
        autoComplete="off"
      />
      <Button
        mt="xl"
        mb="md"
        type="submit"
        sx={{ fontSize: theme.fontSizes.md }}
        loading={loginForm.isSubmitting}
        fullWidth
      >
        {t('general.login')}
      </Button>
    </form>
  )
}
