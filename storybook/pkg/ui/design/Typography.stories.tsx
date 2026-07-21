import type { Meta, StoryObj } from '@storybook/react-vite'
import {
  Body1,
  Body2,
  ButtonText,
  Caption,
  H1,
  H2,
  H3,
  H4,
  H5,
  H6,
  Overline,
  Subtitle1,
  Subtitle2,
} from '@/pkg/ui/ts/design/typography'

const meta = {
  title: 'UI/Design/Typography',
  parameters: {
    layout: 'padded',
  },
} satisfies Meta

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  render: () => (
    <div className="flex flex-col gap-4">
      <H1>h1. Heading</H1>
      <H2>h2. Heading</H2>
      <H3>h3. Heading</H3>
      <H4>h4. Heading</H4>
      <H5>h5. Heading</H5>
      <H6>h6. Heading</H6>
      <Subtitle1>subtitle1. Lorem ipsum dolor sit amet</Subtitle1>
      <Subtitle2>subtitle2. Lorem ipsum dolor sit amet</Subtitle2>
      <Body1>body1. Lorem ipsum dolor sit amet</Body1>
      <Body2>body2. Lorem ipsum dolor sit amet</Body2>
      <ButtonText>button text</ButtonText>
      <Caption>caption text</Caption>
      <Overline>overline text</Overline>
    </div>
  ),
}
