import { useEffect } from 'react'
import { Route, Routes, useNavigate } from 'react-router-dom'
import { Events } from '@wailsio/runtime'
import GordleView, { GordleCell, GordleCellState } from  '@/cmd/gordle_app/frontend/src/pages/gordle/GordleView'
import MainMenuView from '@/cmd/gordle_app/frontend/src/pages/main_menu/MainMenuView'

const initialGordleCells: GordleCell[][] = Array.from({ length: 6 }, () =>
  Array.from({ length: 5 }, () => ({ letter: '', state: GordleCellState.Unguessed })),
)

function App() {
  const navigate = useNavigate()

  useEffect(() => {
    Events.On('gordle_app:navigate:path', (event) => navigate(event.data))
  }, [])

  return (
    <Routes>
      <Route
        path="/"
        element={
          <MainMenuView />
        }
      />
      <Route path="/gordle" element={<GordleView cells={initialGordleCells} />} />
    </Routes>
  )
}

export default App
