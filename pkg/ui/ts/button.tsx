import Button from '@mui/material/Button'
import React from 'react'

export type ButtonProps = {
  title: string
  onClick: () => void
}

export const PrimaryButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <Button color='primary' variant='contained' onClick={onClick}>{title}</Button>
}

export const SecondaryButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <Button color='secondary' variant='contained' onClick={onClick}>{title}</Button>
}

export const ErrorButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <Button color='error' variant='contained' onClick={onClick}>{title}</Button>
}

export const InfoButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <Button color='info' variant='contained' onClick={onClick}>{title}</Button>
}

export const SuccessButton = ({ title, onClick }: ButtonProps): React.JSX.Element => {
  return <Button color='success' variant='contained' onClick={onClick}>{title}</Button>
}
