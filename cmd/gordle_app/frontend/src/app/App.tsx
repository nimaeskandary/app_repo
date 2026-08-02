import { useMemo } from 'react'
import { Route, Routes, useNavigate } from 'react-router-dom'
import {
  DependenciesProvider,
  type Dependencies,
} from '@app_repo/pkg_gordle/app/Dependencies'
import GordlePage from '@app_repo/pkg_gordle/pages/gordle/GordlePage'
import MainMenuPage from '@app_repo/pkg_gordle/pages/main_menu/MainMenuPage'
import Layout from './Layout'

function App() {
  const navigate = useNavigate()
  const dependencies = useMemo<Dependencies>(() => ({
    navigator: {
      goToGame: () => navigate('/gordle'),
    },
  }), [navigate])

  return (
    <Layout>
      <DependenciesProvider value={dependencies}>
        <Routes>
          <Route path="/" element={<MainMenuPage />} />
          <Route path="/gordle" element={<GordlePage />} />
        </Routes>
      </DependenciesProvider>
    </Layout>
  )
}

export default App
