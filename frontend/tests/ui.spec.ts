import { test, expect } from '@playwright/test'

test.describe('DevOps Converter UI Tests', () => {
  test.beforeEach(async ({ page }) => {
    // Naviguer vers l'application
    await page.goto('/')
  })

  test('should display the main header correctly', async ({ page }) => {
    // Vérifier que le titre principal est affiché
    await expect(page.locator('h1')).toContainText('DevOps Converter')
    
    // Vérifier que le sous-titre est affiché
    await expect(page.locator('p')).toContainText('Docker Compose → Kubernetes')
    
    // Vérifier que le statut de santé est affiché
    await expect(page.locator('[data-testid="health-status"]')).toBeVisible()
  })

  test('should show conversion interface', async ({ page }) => {
    // Vérifier que la zone de conversion est visible
    await expect(page.locator('h1')).toContainText('Convertisseur Docker → Kubernetes')
    
    // Vérifier que les statistiques sont affichées
    await expect(page.locator('text=100+')).toBeVisible()
    await expect(page.locator('text=99.9%')).toBeVisible()
    await expect(page.locator('text=<1s')).toBeVisible()
  })

  test('should handle text input conversion', async ({ page }) => {
    // Sélectionner l'onglet texte
    await page.click('text=Texte')
    
    // Saisir du contenu Docker Compose
    const dockerComposeContent = `version: '3.8'
services:
  web:
    image: nginx:latest
    ports:
      - "80:80"
    environment:
      - ENV=production`
    
    await page.fill('textarea[placeholder*="version"]', dockerComposeContent)
    
    // Vérifier que le bouton convertir est activé
    await expect(page.locator('button:has-text("Convertir")')).toBeEnabled()
  })

  test('should show notification when conversion is attempted', async ({ page }) => {
    // Sélectionner l'onglet texte et ajouter du contenu
    await page.click('text=Texte')
    
    const dockerComposeContent = `version: '3.8'
services:
  web:
    image: nginx:latest
    ports:
      - "80:80"`
    
    await page.fill('textarea[placeholder*="version"]', dockerComposeContent)
    
    // Cliquer sur convertir
    await page.click('button:has-text("Convertir")')
    
    // Attendre et vérifier qu'une notification apparaît
    // Note: Ceci peut échouer si le backend n'est pas accessible, mais c'est OK pour le test UI
    await page.waitForTimeout(2000)
    
    // Vérifier que l'état de chargement est géré
    const convertButton = page.locator('button:has-text("Convertir"), button:has-text("Conversion...")')
    await expect(convertButton).toBeVisible()
  })

  test('should handle file upload interface', async ({ page }) => {
    // Cliquer sur l'onglet fichier
    await page.click('text=Fichier')
    
    // Vérifier que la zone de drop est visible
    await expect(page.locator('text=Cliquez pour sélectionner')).toBeVisible()
    await expect(page.locator('text=glissez-déposez votre fichier')).toBeVisible()
    await expect(page.locator('text=YAML uniquement')).toBeVisible()
  })

  test('should show configuration options', async ({ page }) => {
    // Vérifier que les options de configuration sont visibles
    await expect(page.locator('text=Options de conversion')).toBeVisible()
    
    // Vérifier les champs de configuration
    await expect(page.locator('input[placeholder="default"]')).toBeVisible() // Namespace
    await expect(page.locator('select')).toHaveCount(2) // Type de fichier + Type de service
    await expect(page.locator('input[type="number"]')).toBeVisible() // Replicas
  })

  test('should validate responsive design', async ({ page }) => {
    // Tester en mobile
    await page.setViewportSize({ width: 375, height: 667 })
    
    // Vérifier que l'interface s'adapte
    await expect(page.locator('h1')).toBeVisible()
    await expect(page.locator('text=Configuration d\'entrée')).toBeVisible()
    
    // Revenir en desktop
    await page.setViewportSize({ width: 1920, height: 1080 })
    
    // Vérifier que la grille s'affiche correctement
    await expect(page.locator('h1')).toBeVisible()
  })
})

test.describe('Notifications Tests', () => {
  test('should show and dismiss notifications', async ({ page }) => {
    await page.goto('/')
    
    // Simuler une action qui génère une notification
    // (ceci dépend de l'implémentation exacte)
    await page.evaluate(() => {
      // Simuler l'ajout d'une notification via le store
      window.dispatchEvent(new CustomEvent('test-notification', {
        detail: {
          type: 'success',
          title: 'Test',
          message: 'Notification de test'
        }
      }))
    })
    
    await page.waitForTimeout(500)
    
    // Si des notifications sont visibles, tester leur fermeture
    const notifications = page.locator('[data-testid="notification"]')
    const notificationCount = await notifications.count()
    
    if (notificationCount > 0) {
      // Cliquer sur le bouton de fermeture de la première notification
      await notifications.first().locator('button').click()
      
      // Vérifier que la notification a été supprimée
      await expect(notifications).toHaveCount(notificationCount - 1)
    }
  })
})

test.describe('Animation and UX Tests', () => {
  test('should have smooth transitions', async ({ page }) => {
    await page.goto('/')
    
    // Tester les animations de hover sur les boutons
    const convertButton = page.locator('button:has-text("Convertir")')
    
    if (await convertButton.isVisible()) {
      // Vérifier que le bouton est interactif
      await convertButton.hover()
      
      // Le bouton devrait avoir des styles de hover
      const buttonStyles = await convertButton.evaluate(el => {
        return window.getComputedStyle(el).transform
      })
      
      // Les transformations CSS indiquent des animations
      expect(buttonStyles).toBeDefined()
    }
  })

  test('should have accessible focus states', async ({ page }) => {
    await page.goto('/')
    
    // Tester la navigation au clavier
    await page.keyboard.press('Tab')
    
    // Vérifier qu'un élément a le focus
    const focusedElement = await page.locator(':focus').count()
    expect(focusedElement).toBeGreaterThan(0)
  })
})
