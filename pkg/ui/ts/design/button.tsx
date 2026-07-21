import MuiButton from '@mui/material/Button'
import React from 'react'

export type ButtonProps = {
  title: string
  onClick: () => void
}

export const PrimaryButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <MuiButton color='primary' variant='contained' onClick={onClick}>{title}</MuiButton>
}

export const SecondaryButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <MuiButton color='secondary' variant='contained' onClick={onClick}>{title}</MuiButton>
}

export const ErrorButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <MuiButton color='error' variant='contained' onClick={onClick}>{title}</MuiButton>
}

export const InfoButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <MuiButton color='info' variant='contained' onClick={onClick}>{title}</MuiButton>
}

export const SuccessButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <MuiButton color='success' variant='contained' onClick={onClick}>{title}</MuiButton>
}
