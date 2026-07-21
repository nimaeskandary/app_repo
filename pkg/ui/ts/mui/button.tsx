import MuiButton, { type ButtonProps as MuiButtonProps } from '@mui/material/Button'

export type ButtonProps = MuiButtonProps

export const PrimaryButton = (props: ButtonProps) => {
  return <MuiButton color="primary" variant="contained" {...props} />
}

export const SecondaryButton = (props: ButtonProps) => {
  return <MuiButton color="secondary" variant="contained" {...props} />
}

export const ErrorButton = (props: ButtonProps) => {
  return <MuiButton color="error" variant="contained" {...props} />
}

export const InfoButton = (props: ButtonProps) => {
  return <MuiButton color="info" variant="contained" {...props} />
}

export const SuccessButton = (props: ButtonProps) => {
  return <MuiButton color="success" variant="contained" {...props} />
}
