import Button from '@mui/material/Button'
import React from 'react'

export type PrimaryButtonProps = {
  title: string
  onClick: () => void
}

export const PrimaryButton = ({ title, onClick }: PrimaryButtonProps): React.JSX.Element => {
  return <Button onClick={onClick}>{title}</Button>
}
