import { test, expect } from '@playwright/test'

test.describe('Démonstration Playwright - Interface DevOps Converter', () => {
  test('démo complète de l\'interface utilisateur', async ({ page }) => {
    console.log('🚀 Démarrage de la démonstration Playwright')
    
    // 1. Navigation vers l'application
    await page.goto('/')
    console.log('✅ Page chargée')
    
    // 2. Vérification de l'interface principale
    await expect(page.locator('h1').first()).toContainText('DevOps Converter')
    console.log('✅ Header vérifié')
    
    // 3. Vérification du statut de santé
    const healthStatus = page.locator('[data-testid="health-status"]')
    await expect(healthStatus).toBeVisible()
    console.log('✅ Statut de santé affiché')
    
    // 4. Test de la page de conversion
    await expect(page.locator('text=Convertisseur Docker → Kubernetes')).toBeVisible()
    console.log('✅ Page de conversion affichée')
    
    // 5. Test des statistiques
    await expect(page.locator('text=100+')).toBeVisible()
    await expect(page.locator('text=99.9%')).toBeVisible()
    await expect(page.locator('text=<1s')).toBeVisible()
    console.log('✅ Statistiques affichées')
    
    // 6. Test de saisie de texte
    await page.click('text=Texte')
    console.log('✅ Onglet Texte sélectionné')
    
    const textarea = page.locator('textarea')
    const dockerContent = `version: '3.8'
services:
  nginx:
    image: nginx:latest
    ports:
      - "80:80"
    environment:
      - ENV=production
  redis:
    image: redis:alpine
    ports:
      - "6379:6379"
    
volumes:
  data:`
    
    await textarea.fill(dockerContent)
    console.log('✅ Contenu Docker Compose saisi')
    
    // 7. Configuration des options
    await page.fill('input[placeholder="default"]', 'demo-namespace')
    console.log('✅ Namespace configuré')
    
    await page.selectOption('select', 'LoadBalancer')
    console.log('✅ Type de service configuré')
    
    await page.fill('input[type="number"]', '3')
    console.log('✅ Nombre de replicas configuré')
    
    // 8. Test des boutons
    const validateButton = page.locator('button:has-text("Valider")')
    const convertButton = page.locator('button:has-text("Convertir")')
    
    await expect(validateButton).toBeEnabled()
    await expect(convertButton).toBeEnabled()
    console.log('✅ Boutons activés')
    
    // 9. Test de validation
    await validateButton.click()
    console.log('✅ Validation tentée')
    
    await page.waitForTimeout(2000)
    
    // 10. Test de conversion
    await convertButton.click()
    console.log('✅ Conversion tentée')
    
    await page.waitForTimeout(3000)
    
    // 11. Test de l'onglet fichier
    await page.click('text=Fichier')
    await expect(page.locator('text=Cliquez pour sélectionner')).toBeVisible()
    console.log('✅ Interface de fichier testée')
    
    // 12. Test responsive
    await page.setViewportSize({ width: 375, height: 667 })
    await expect(page.locator('h1').first()).toBeVisible()
    console.log('✅ Design responsive testé (mobile)')
    
    await page.setViewportSize({ width: 1920, height: 1080 })
    await expect(page.locator('h1').first()).toBeVisible()
    console.log('✅ Design responsive testé (desktop)')
    
    // 13. Capture d'écran finale
    await page.screenshot({ path: 'tests/screenshots/demo-final.png', fullPage: true })
    console.log('✅ Capture d\'écran sauvegardée')
    
    console.log('🎉 Démonstration Playwright terminée avec succès!')
  })

  test('test de performance et chargement', async ({ page }) => {
    const startTime = Date.now()
    
    await page.goto('/')
    
    // Attendre que tous les éléments critiques soient chargés
    await expect(page.locator('h1').first()).toBeVisible()
    await expect(page.locator('[data-testid="health-status"]')).toBeVisible()
    await expect(page.locator('textarea')).toBeVisible()
    
    const loadTime = Date.now() - startTime
    console.log(`⏱️  Temps de chargement total: ${loadTime}ms`)
    
    // Vérifier que la page se charge rapidement
    expect(loadTime).toBeLessThan(10000) // 10 secondes max
    
    if (loadTime < 2000) {
      console.log('🚀 Performance excellente!')
    } else if (loadTime < 5000) {
      console.log('⚡ Performance correcte')
    } else {
      console.log('⚠️  Performance à améliorer')
    }
  })

  test('test d\'accessibilité de base', async ({ page }) => {
    await page.goto('/')
    
    // Test de navigation au clavier
    await page.keyboard.press('Tab')
    await page.keyboard.press('Tab')
    await page.keyboard.press('Tab')
    
    const focusedElement = page.locator(':focus')
    await expect(focusedElement).toBeVisible()
    console.log('✅ Navigation clavier fonctionnelle')
    
    // Vérifier les contrastes de base (éléments visibles)
    const buttons = page.locator('button')
    const buttonCount = await buttons.count()
    expect(buttonCount).toBeGreaterThan(0)
    console.log(`✅ ${buttonCount} boutons détectés`)
    
    // Vérifier la structure sémantique
    await expect(page.locator('main')).toBeVisible()
    await expect(page.locator('header')).toBeVisible()
    console.log('✅ Structure sémantique correcte')
  })
})
