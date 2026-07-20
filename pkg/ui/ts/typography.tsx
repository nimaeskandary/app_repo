import MuiTypography from '@mui/material/Typography'
import React from 'react'

export type TypographyProps = {
  children: React.ReactNode
}

export const H1 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="h1">{children}</MuiTypography>
}

export const H2 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="h2">{children}</MuiTypography>
}

export const H3 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="h3">{children}</MuiTypography>
}

export const H4 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="h4">{children}</MuiTypography>
}

export const H5 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="h5">{children}</MuiTypography>
}

export const H6 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="h6">{children}</MuiTypography>
}

export const Subtitle1 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="subtitle1">{children}</MuiTypography>
}

export const Subtitle2 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="subtitle2">{children}</MuiTypography>
}

export const Body1 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="body1">{children}</MuiTypography>
}

export const Body2 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="body2">{children}</MuiTypography>
}

export const ButtonText = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="button">{children}</MuiTypography>
}

export const Caption = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="caption">{children}</MuiTypography>
}

export const Overline = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="overline">{children}</MuiTypography>
}
