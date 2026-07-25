import { useMemo } from 'react'
import { Route, Routes, useNavigate } from 'react-router-dom'
import {
  GordleDependenciesProvider,
  type GordleDependencies,
} from '@app_repo/pkg_gordle_frontend/di/GordleDependencies'
import GordlePage from '@app_repo/pkg_gordle_frontend/views/gordle/GordlePage'
import MainMenuPage from '@app_repo/pkg_gordle_frontend/views/main_menu/MainMenuPage'

function App() {
  const navigate = useNavigate()
  const dependencies = useMemo<GordleDependencies>(() => ({
    navigator: {
      goToGame: () => navigate('/gordle'),
    },
  }), [navigate])

  return (
    <GordleDependenciesProvider value={dependencies}>
      <Routes>
        <Route path="/" element={<MainMenuPage />} />
        <Route path="/gordle" element={<GordlePage />} />
      </Routes>
    </GordleDependenciesProvider>
  )
}

export default App
