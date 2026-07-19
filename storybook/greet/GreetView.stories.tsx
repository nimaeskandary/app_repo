import type { Meta, StoryObj } from '@storybook/react-vite'
import { useArgs } from 'storybook/preview-api'
import GreetView from '../../app/wails_app/frontend/src/GreetView'

const meta = {
  title: 'Greet/GreetView',
  component: GreetView,
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
    time: new Date().toUTCString(),
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
