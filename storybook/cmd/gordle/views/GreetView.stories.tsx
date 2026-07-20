import type { Meta, StoryObj } from '@storybook/react-vite'
import { useArgs } from 'storybook/preview-api'
import GreetView from '../../../../cmd/gordle_app/frontend/src/GreetView'
import GordlePreview from '../GordlePreview'

const meta = {
  title: 'Apps/Gordle/Views/GreetView',
  component: GreetView,
  decorators: [
    (Story) => (
      <GordlePreview>
        <Story />
      </GordlePreview>
    ),
  ],
  parameters: {
    layout: 'fullscreen',
  },
  argTypes: {
    name: { control: 'text' },
    titleName: { control: 'text' },
    time: { control: 'text' },
    toastMessage: { control: 'text' },
    isToastVisible: { control: 'boolean' },
    onNameChange: { table: { disable: true } },
    onGreet: { table: { disable: true } },
  },
} satisfies Meta<typeof GreetView>

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    name: '',
    titleName: 'React',
    time: '02:35:06',
    toastMessage: '',
    isToastVisible: false,
    onNameChange: () => {},
    onGreet: () => {},
  },
  render: (args) => {
    const [, updateArgs] = useArgs()
    const greet = () => {
      const nextName = args.name || 'anonymous'
      updateArgs({
        titleName: nextName,
        toastMessage: `Hello ${nextName}`,
        isToastVisible: true,
      })
    }

    return (
      <GreetView
        {...args}
        onNameChange={(name) => updateArgs({ name })}
        onGreet={greet}
      />
    )
  },
}
