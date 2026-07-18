import { useEffect, useRef, useState } from 'react'
import { Events, WML } from '@wailsio/runtime'
import { GreetService } from '@/bases/wails_app/frontend/bindings/github.com/nimaeskandary/wails3-react-polylith/bases/wails_app/app/bridge'
import GreetView from './GreetView'

function App() {
  const [name, setName] = useState<string>('')
  const [titleName, setTitleName] = useState<string>('React')
  const [time, setTime] = useState<string>('Listening for Time event...')
  const [toastMessage, setToastMessage] = useState<string>('')
  const [isToastVisible, setIsToastVisible] = useState<boolean>(false)
  const toastTimer = useRef<ReturnType<typeof setTimeout> | undefined>(undefined)

  // Pop the toast with the message Go returned, then auto-dismiss it.
  const showToast = (message: string) => {
    setToastMessage(message)
    setIsToastVisible(true)
    clearTimeout(toastTimer.current)
    toastTimer.current = setTimeout(() => {
      setIsToastVisible(false)
    }, 4000)
  }

  const doGreet = () => {
    const n = name || 'anonymous'
    setTitleName(n)
    GreetService.Greet(n).then(showToast).catch(console.error)
  }

  useEffect(() => {
    Events.On('time', (timeValue: any) => {
      // On a narrow screen the full RFC1123 stamp is too wide for the footer, so
      // show just the clock time there (matching the CSS breakpoint).
      const full = timeValue.data
      const compact = (full.match(/\d{1,2}:\d{2}:\d{2}/) || [full])[0]
      setTime(window.matchMedia('(max-width: 640px)').matches ? compact : full)
    })
    // Reload WML so it picks up the wml tags
    WML.Reload()
  }, [])

  return (
    <GreetView
      name={name}
      titleName={titleName}
      time={time}
      toastMessage={toastMessage}
      isToastVisible={isToastVisible}
      onNameChange={setName}
      onGreet={doGreet}
    />
  )
}

export default App
