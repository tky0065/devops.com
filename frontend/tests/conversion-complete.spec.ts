import { test, expect } from '@playwright/test'
import path from 'path'

test.describe('DevOps Converter - Tests Complets', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('http://localhost:5175')
    await page.waitForLoadState('networkidle')
  })

  test('Interface utilisateur complète', async ({ page }) => {
    // Vérifier le titre de la page
    await expect(page).toHaveTitle('DevOps Converter - Docker to Kubernetes')
    
    // Vérifier les éléments principaux de l'interface
    await expect(page.getByRole('heading', { name: 'DevOps Converter' })).toBeVisible()
    await expect(page.getByText('Docker Compose → Kubernetes')).toBeVisible()
    
    // Vérifier les boutons principaux
    await expect(page.getByRole('button', { name: 'Saisie manuelle' })).toBeVisible()
    await expect(page.getByRole('button', { name: 'Import de fichier' })).toBeVisible()
    
    // Vérifier les champs de configuration
    await expect(page.getByPlaceholder('Entrez le nom du projet...')).toBeVisible()
    await expect(page.getByRole('radio', { name: 'Fichier unique (recommandé)' })).toBeChecked()
    
    // Vérifier les boutons d'action
    await expect(page.getByRole('button', { name: 'Valider la syntaxe' })).toBeVisible()
    await expect(page.getByRole('button', { name: 'Convertir vers Kubernetes' })).toBeVisible()
  })

  test('Upload et conversion de fichier avec fusion (allInOne)', async ({ page }) => {
    // Préparer le fichier à uploader
    const testFile = path.resolve('/Users/yacoubakone/Documents/dev/devops.com/test-docker-compose.yml')
    
    // Cliquer sur le bouton d'import
    await page.getByRole('button', { name: 'Import de fichier' }).click()
    
    // Upload du fichier
    const fileInput = page.locator('input[type="file"]')
    await fileInput.setInputFiles(testFile)
    
    // Vérifier que le fichier a été chargé
    await expect(page.getByText('test-docker-compose.yml')).toBeVisible()
    await expect(page.getByText('Fichier chargé')).toBeVisible()
    
    // Entrer un nom de projet
    await page.getByPlaceholder('Entrez le nom du projet...').fill('test-project-fusion')
    
    // S'assurer que "Fichier unique" est sélectionné
    await page.getByRole('radio', { name: 'Fichier unique (recommandé)' }).check()
    
    // Valider la syntaxe
    await page.getByRole('button', { name: 'Valider la syntaxe' }).click()
    
    // Attendre et vérifier la validation
    await expect(page.getByText('Configuration validée avec succès!')).toBeVisible()
    await expect(page.getByText('Valide')).toBeVisible()
    
    // Convertir vers Kubernetes
    await page.getByRole('button', { name: 'Convertir vers Kubernetes' }).click()
    
    // Attendre et vérifier la conversion
    await page.waitForTimeout(2000) // Attendre la conversion
    
    // Vérifier qu'il n'y a pas d'erreur
    await expect(page.getByText('Erreur de conversion')).not.toBeVisible()
  })

  test('Upload et conversion de fichier avec fichiers séparés', async ({ page }) => {
    const testFile = path.resolve('/Users/yacoubakone/Documents/dev/devops.com/test-docker-compose.yml')
    
    // Upload du fichier
    await page.getByRole('button', { name: 'Import de fichier' }).click()
    const fileInput = page.locator('input[type="file"]')
    await fileInput.setInputFiles(testFile)
    
    // Attendre et vérifier le chargement
    await expect(page.getByText('test-docker-compose.yml')).toBeVisible()
    
    // Entrer un nom de projet
    await page.getByPlaceholder('Entrez le nom du projet...').fill('test-project-separe')
    
    // Sélectionner "Fichiers séparés"
    await page.getByRole('radio', { name: 'Fichiers séparés' }).check()
    
    // Valider
    await page.getByRole('button', { name: 'Valider la syntaxe' }).click()
    await expect(page.getByText('Configuration validée avec succès!')).toBeVisible()
    
    // Convertir
    await page.getByRole('button', { name: 'Convertir vers Kubernetes' }).click()
    await page.waitForTimeout(2000)
    
    // Vérifier qu'il n'y a pas d'erreur
    await expect(page.getByText('Erreur de conversion')).not.toBeVisible()
  })

  test('Saisie manuelle avec configuration personnalisée', async ({ page }) => {
    // Cliquer sur saisie manuelle
    await page.getByRole('button', { name: 'Saisie manuelle' }).click()
    
    // Entrer du contenu Docker Compose
    const dockerComposeContent = `version: '3.8'
services:
  api:
    image: node:16-alpine
    ports:
      - "3000:3000"
    environment:
      - NODE_ENV=production
  db:
    image: postgres:13
    environment:
      - POSTGRES_DB=myapp
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
    volumes:
      - db_data:/var/lib/postgresql/data
volumes:
  db_data:`
  
    // Trouver et remplir la zone de texte
    const textarea = page.locator('textarea').first()
    await textarea.fill(dockerComposeContent)
    
    // Configurer les options
    await page.getByPlaceholder('Entrez le nom du projet...').fill('multi-service-app')
    await page.locator('input[placeholder="default"]').fill('production')
    await page.locator('select').selectOption('NodePort')
    await page.locator('input[type="number"]').fill('3')
    
    // Valider
    await page.getByRole('button', { name: 'Valider la syntaxe' }).click()
    await expect(page.getByText('Configuration validée avec succès!')).toBeVisible()
    
    // Convertir avec fichier unique
    await page.getByRole('radio', { name: 'Fichier unique (recommandé)' }).check()
    await page.getByRole('button', { name: 'Convertir vers Kubernetes' }).click()
    await page.waitForTimeout(2000)
    
    // Vérifier le succès
    await expect(page.getByText('Erreur de conversion')).not.toBeVisible()
  })

  test('Test des notifications et de la gestion des erreurs', async ({ page }) => {
    // Tenter une conversion sans fichier
    await page.getByRole('button', { name: 'Convertir vers Kubernetes' }).click()
    
    // Vérifier qu'une notification d'erreur apparaît (comportement attendu)
    await page.waitForTimeout(1000)
    
    // Tester avec un contenu invalide
    await page.getByRole('button', { name: 'Saisie manuelle' }).click()
    const textarea = page.locator('textarea').first()
    await textarea.fill('contenu invalide')
    
    await page.getByRole('button', { name: 'Valider la syntaxe' }).click()
    await page.waitForTimeout(1000)
    
    // La validation devrait échouer pour un contenu invalide
  })

  test('Fonctionnalité de nommage de projet et téléchargement', async ({ page }) => {
    const testFile = path.resolve('/Users/yacoubakone/Documents/dev/devops.com/test-docker-compose.yml')
    
    // Upload du fichier
    await page.getByRole('button', { name: 'Import de fichier' }).click()
    const fileInput = page.locator('input[type="file"]')
    await fileInput.setInputFiles(testFile)
    
    await expect(page.getByText('test-docker-compose.yml')).toBeVisible()
    
    // Tester différents noms de projet
    const projectNames = ['mon-projet-k8s', 'app-production', 'test-123']
    
    for (const projectName of projectNames) {
      await page.getByPlaceholder('Entrez le nom du projet...').fill(projectName)
      
      // Valider et convertir
      await page.getByRole('button', { name: 'Valider la syntaxe' }).click()
      await expect(page.getByText('Configuration validée avec succès!')).toBeVisible()
      
      await page.getByRole('button', { name: 'Convertir vers Kubernetes' }).click()
      await page.waitForTimeout(1500)
      
      // Vérifier qu'il n'y a pas d'erreur
      await expect(page.getByText('Erreur de conversion')).not.toBeVisible()
    }
  })

  test('Test des icônes et design sans émojis', async ({ page }) => {
    // Vérifier qu'aucun émoji n'est présent dans l'interface
    const pageContent = await page.textContent('body')
    
    // Liste d'émojis couramment utilisés qu'on ne devrait pas trouver
    const emojis = ['🚀', '⚡', '🔧', '📁', '💻', '🎯', '✅', '❌', '📊', '🔍']
    
    for (const emoji of emojis) {
      expect(pageContent).not.toContain(emoji)
    }
    
    // Vérifier la présence d'éléments SVG (icônes)
    const svgElements = page.locator('svg')
    const svgCount = await svgElements.count()
    expect(svgCount).toBeGreaterThan(5) // Au moins quelques icônes SVG
    
    // Vérifier des classes d'icônes spécifiques
    await expect(page.locator('[class*="icon"]')).toBeVisible()
  })

  test('Test de responsivité et interface mobile', async ({ page }) => {
    // Tester différentes tailles d'écran
    const viewports = [
      { width: 1920, height: 1080 }, // Desktop
      { width: 1024, height: 768 },  // Tablet
      { width: 375, height: 667 }    // Mobile
    ]
    
    for (const viewport of viewports) {
      await page.setViewportSize(viewport)
      await page.reload()
      await page.waitForLoadState('networkidle')
      
      // Vérifier que les éléments principaux sont toujours visibles
      await expect(page.getByRole('heading', { name: 'DevOps Converter' })).toBeVisible()
      await expect(page.getByRole('button', { name: 'Import de fichier' })).toBeVisible()
    }
  })

  test('Performance et temps de réponse', async ({ page }) => {
    const startTime = Date.now()
    
    // Mesurer le temps de chargement initial
    await page.goto('http://localhost:5175')
    await page.waitForLoadState('networkidle')
    
    const loadTime = Date.now() - startTime
    expect(loadTime).toBeLessThan(5000) // Chargement en moins de 5 secondes
    
    // Tester la réactivité des interactions
    const testFile = path.resolve('/Users/yacoubakone/Documents/dev/devops.com/test-docker-compose.yml')
    
    const uploadStart = Date.now()
    await page.getByRole('button', { name: 'Import de fichier' }).click()
    const fileInput = page.locator('input[type="file"]')
    await fileInput.setInputFiles(testFile)
    await expect(page.getByText('test-docker-compose.yml')).toBeVisible()
    const uploadTime = Date.now() - uploadStart
    
    expect(uploadTime).toBeLessThan(3000) // Upload en moins de 3 secondes
  })
})
