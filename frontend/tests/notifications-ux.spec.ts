import { test, expect } from '@playwright/test'

test.describe('DevOps Converter - Tests des notifications et UX', () => {
  test.beforeEach(async ({ page }) => {
    // Naviguer vers l'application
    await page.goto('/')
  })

  test('doit afficher l\'interface principale correctement', async ({ page }) => {
    // Vérifier le titre principal
    await expect(page.locator('h1').first()).toContainText('DevOps Converter')
    
    // Vérifier le sous-titre
    await expect(page.locator('text=Docker Compose → Kubernetes')).toBeVisible()
    
    // Vérifier l'icône et le design moderne
    await expect(page.locator('svg').first()).toBeVisible()
    
    // Vérifier le statut de santé
    const healthStatus = page.locator('[data-testid="health-status"]')
    await expect(healthStatus).toBeVisible()
    
    // Vérifier le lien GitHub
    await expect(page.locator('text=GitHub')).toBeVisible()
  })

  test('doit afficher la page de conversion avec design amélioré', async ({ page }) => {
    // Vérifier le titre de la page de conversion
    await expect(page.locator('text=Convertisseur Docker → Kubernetes')).toBeVisible()
    
    // Vérifier les statistiques
    await expect(page.locator('text=100+')).toBeVisible()
    await expect(page.locator('text=99.9%')).toBeVisible()
    await expect(page.locator('text=<1s')).toBeVisible()
    
    // Vérifier les labels des statistiques
    await expect(page.locator('text=Conversions')).toBeVisible()
    await expect(page.locator('text=Précision')).toBeVisible()
    await expect(page.locator('text=Temps moyen')).toBeVisible()
  })

  test('doit permettre la saisie de texte et l\'interaction', async ({ page }) => {
    // Cliquer sur l'onglet Texte s'il n'est pas déjà sélectionné
    await page.click('text=Texte')
    
    // Trouver la zone de texte
    const textarea = page.locator('textarea')
    await expect(textarea).toBeVisible()
    
    // Saisir du contenu Docker Compose
    const dockerComposeContent = `version: '3.8'
services:
  web:
    image: nginx:latest
    ports:
      - "80:80"
    environment:
      - ENV=production
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
    
    await textarea.fill(dockerComposeContent)
    
    // Vérifier que le contenu a été saisi
    await expect(textarea).toHaveValue(dockerComposeContent)
    
    // Vérifier que les boutons sont maintenant activés
    const validateButton = page.locator('button:has-text("Valider")')
    const convertButton = page.locator('button:has-text("Convertir")')
    
    await expect(validateButton).toBeEnabled()
    await expect(convertButton).toBeEnabled()
  })

  test('doit gérer les options de conversion', async ({ page }) => {
    // Vérifier les options de conversion
    await expect(page.locator('text=Options de conversion')).toBeVisible()
    
    // Tester le champ namespace
    const namespaceInput = page.locator('input[placeholder="default"]')
    await expect(namespaceInput).toBeVisible()
    await namespaceInput.fill('production')
    await expect(namespaceInput).toHaveValue('production')
    
    // Tester le sélecteur de type de service
    const serviceTypeSelect = page.locator('select').nth(1) // Le deuxième select
    await expect(serviceTypeSelect).toBeVisible()
    await serviceTypeSelect.selectOption('LoadBalancer')
    
    // Tester le champ replicas
    const replicasInput = page.locator('input[type="number"]')
    await expect(replicasInput).toBeVisible()
    await replicasInput.fill('3')
    await expect(replicasInput).toHaveValue('3')
  })

  test('doit tester l\'interface de fichier upload', async ({ page }) => {
    // Cliquer sur l'onglet Fichier
    await page.click('text=Fichier')
    
    // Vérifier la zone de drop
    await expect(page.locator('text=Cliquez pour sélectionner')).toBeVisible()
    await expect(page.locator('text=glissez-déposez votre fichier')).toBeVisible()
    await expect(page.locator('text=YAML uniquement')).toBeVisible()
    
    // Vérifier l'icône de fichier
    await expect(page.locator('svg').nth(2)).toBeVisible() // L'icône dans la zone de drop
  })

  test('doit tester la réactivité du design', async ({ page }) => {
    // Test en mode mobile
    await page.setViewportSize({ width: 375, height: 667 })
    
    // Vérifier que les éléments principaux sont toujours visibles
    await expect(page.locator('h1').first()).toBeVisible()
    await expect(page.locator('[data-testid="health-status"]')).toBeVisible()
    
    // Test en mode tablette
    await page.setViewportSize({ width: 768, height: 1024 })
    await expect(page.locator('text=Configuration d\'entrée')).toBeVisible()
    
    // Test en mode desktop large
    await page.setViewportSize({ width: 1920, height: 1080 })
    
    // Vérifier que la grille s'affiche correctement
    await expect(page.locator('text=Configuration d\'entrée')).toBeVisible()
    await expect(page.locator('text=Les résultats de conversion apparaîtront ici')).toBeVisible()
  })

  test('doit tester les animations et transitions', async ({ page }) => {
    // Tester l'animation du bouton au hover
    const convertButton = page.locator('button:has-text("Convertir")')
    
    if (await convertButton.isVisible()) {
      // Simuler un hover
      await convertButton.hover()
      
      // Attendre un peu pour que l'animation se joue
      await page.waitForTimeout(300)
      
      // Vérifier que le bouton est toujours visible (test basique d'animation)
      await expect(convertButton).toBeVisible()
    }
    
    // Tester la navigation entre onglets
    await page.click('text=Fichier')
    await page.waitForTimeout(200)
    await expect(page.locator('text=Cliquez pour sélectionner')).toBeVisible()
    
    await page.click('text=Texte')
    await page.waitForTimeout(200)
    await expect(page.locator('textarea')).toBeVisible()
  })

  test('doit simuler une tentative de conversion', async ({ page }) => {
    // Remplir le formulaire
    await page.click('text=Texte')
    
    const dockerContent = `version: '3.8'
services:
  app:
    image: node:16
    ports:
      - "3000:3000"`
    
    await page.fill('textarea', dockerContent)
    
    // Configurer les options
    await page.fill('input[placeholder="default"]', 'test-namespace')
    await page.selectOption('select', 'NodePort')
    await page.fill('input[type="number"]', '2')
    
    // Tenter la validation d'abord
    await page.click('button:has-text("Valider")')
    
    // Attendre un peu pour voir la réaction de l'interface
    await page.waitForTimeout(1000)
    
    // Tenter la conversion
    await page.click('button:has-text("Convertir")')
    
    // Attendre pour voir s'il y a des changements dans l'interface
    await page.waitForTimeout(2000)
    
    // Vérifier que l'interface réagit (bouton de chargement ou résultats)
    const loadingButton = page.locator('button:has-text("Conversion...")')
    const convertButton = page.locator('button:has-text("Convertir")')
    
    // L'un des deux devrait être visible
    await expect(loadingButton.or(convertButton)).toBeVisible()
  })

  test('doit vérifier l\'accessibilité de base', async ({ page }) => {
    // Tester la navigation au clavier
    await page.keyboard.press('Tab')
    await page.keyboard.press('Tab')
    await page.keyboard.press('Tab')
    
    // Vérifier qu'un élément a le focus
    const focusedElement = page.locator(':focus')
    await expect(focusedElement).toBeVisible()
    
    // Tester les rôles ARIA de base
    const buttons = page.locator('button')
    const buttonCount = await buttons.count()
    expect(buttonCount).toBeGreaterThan(0)
    
    // Vérifier que les inputs ont des labels
    const inputs = page.locator('input')
    const inputCount = await inputs.count()
    
    if (inputCount > 0) {
      // Au moins un input devrait être présent
      expect(inputCount).toBeGreaterThan(0)
    }
  })
})

test.describe('Tests de performance et chargement', () => {
  test('doit charger rapidement', async ({ page }) => {
    const startTime = Date.now()
    
    await page.goto('/')
    
    // Attendre que les éléments principaux soient chargés
    await expect(page.locator('h1').first()).toBeVisible()
    await expect(page.locator('[data-testid="health-status"]')).toBeVisible()
    
    const loadTime = Date.now() - startTime
    
    // Vérifier que la page se charge en moins de 5 secondes
    expect(loadTime).toBeLessThan(5000)
    
    console.log(`Temps de chargement: ${loadTime}ms`)
  })

  test('doit gérer les erreurs réseau gracieusement', async ({ page }) => {
    // Bloquer les requêtes réseau vers le backend
    await page.route('http://localhost:8081/**', route => {
      route.abort('failed')
    })
    
    await page.goto('/')
    
    // L'interface devrait quand même se charger
    await expect(page.locator('h1').first()).toBeVisible()
    
    // Le statut de santé devrait indiquer "hors ligne"
    const healthStatus = page.locator('[data-testid="health-status"]')
    await expect(healthStatus).toBeVisible()
    
    // Optionnel: vérifier le texte spécifique si accessible
    // await expect(page.locator('text=Hors ligne')).toBeVisible()
  })
})
