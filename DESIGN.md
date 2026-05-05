# UI/UX Design Guide - GoSecureTransfer

## Design Philosophy

GoSecureTransfer features a **minimalist light theme** design focused on:

- **Clarity**: Clear information hierarchy and intuitive navigation
- **Simplicity**: Minimal visual elements, maximum functionality
- **Accessibility**: WCAG 2.1 compliant with focus states and color contrast
- **Performance**: Lightweight CSS, no heavy dependencies
- **Responsiveness**: Mobile-first approach, works on all devices

## Design System

### Color Palette

| Role | Color | Hex | Usage |
|------|-------|-----|-------|
| Primary | Blue | `#0066cc` | Buttons, links, active states |
| Primary Hover | Dark Blue | `#0052a3` | Button hover states |
| Background | Light Gray | `#f5f5f5` | Page background |
| Surface | White | `#ffffff` | Cards, containers |
| Text Primary | Dark Gray | `#1a1a1a` | Headings, body text |
| Text Secondary | Medium Gray | `#666` | Descriptions, metadata |
| Text Tertiary | Light Gray | `#999` | Timestamps, hints |
| Border | Light Gray | `#e0e0e0` | Dividers, input borders |
| Success | Green | `#0c0` | Success messages |
| Error | Red | `#c00` | Error messages |

### Typography

| Element | Font | Size | Weight | Usage |
|---------|------|------|--------|-------|
| Heading 1 | System | 28px | 600 | Page titles |
| Heading 2 | System | 24px | 600 | Section titles |
| Heading 3 | System | 18px | 600 | Subsection titles |
| Body | System | 14px | 400 | Regular text |
| Small | System | 13px | 400 | Secondary text |
| Label | System | 14px | 500 | Form labels |

**Font Stack:**
```css
-apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",
"Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
sans-serif
```

### Spacing Scale

```
4px   - Extra small (borders, padding in small elements)
8px   - Small (component padding)
12px  - Medium (default padding)
16px  - Large (section padding)
20px  - Extra large (major section spacing)
24px  - XXL (form gaps)
32px  - Huge (page padding)
```

### Border Radius

```
4px   - Small elements (buttons, inputs)
6px   - Medium elements (cards, small components)
8px   - Large elements (major containers)
12px  - Extra large (modals, overlays)
```

### Shadows

```css
/* Small elevation */
box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);

/* Medium elevation */
box-shadow: 0 4px 12px rgba(0, 102, 204, 0.2);

/* Large elevation */
box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);
```

## Component Specifications

### Buttons

#### Primary Button
- **Background:** `#0066cc`
- **Text:** White, 14px, 600 weight
- **Padding:** 12px 16px
- **Border Radius:** 8px
- **Hover:** `#0052a3`, raised by 2px
- **Active:** No transform
- **Disabled:** Opacity 0.5

#### Secondary Button
- **Background:** `#f0f0f0`
- **Text:** `#333`, 14px, 600 weight
- **Border:** 1px `#e0e0e0`
- **Padding:** 8px 16px
- **Border Radius:** 6px
- **Hover:** `#e0e0e0` background

### Form Inputs

#### Text Input
- **Background:** `#fafafa` (default), `#ffffff` (focused)
- **Border:** 1px `#e0e0e0` (default), `#0066cc` (focused)
- **Border Radius:** 8px
- **Padding:** 12px 14px
- **Font Size:** 14px
- **Focus Shadow:** `0 0 0 3px rgba(0, 102, 204, 0.1)`
- **Placeholder Color:** `#ccc`

#### File Upload Area
- **Border:** 2px dashed `#d0d0d0`
- **Border Radius:** 8px
- **Padding:** 40px 20px
- **Background:** `#fafafa` (default), `#f0f7ff` (hover/drag)
- **Hover Border:** `#0066cc`

### Cards

#### Dashboard Card
- **Background:** White
- **Border Radius:** 12px
- **Padding:** 24px
- **Border:** None
- **Shadow:** `0 2px 12px rgba(0, 0, 0, 0.08)`
- **Hover:** Subtle shadow increase

### Messages

#### Success Message
- **Background:** `#efe`
- **Text:** `#3c3`
- **Border Left:** 4px solid `#3c3`
- **Padding:** 12px 14px
- **Border Radius:** 8px

#### Error Message
- **Background:** `#fee`
- **Text:** `#c33`
- **Border Left:** 4px solid `#c33`
- **Padding:** 12px 14px
- **Border Radius:** 8px

## Page Layouts

### Login / Register Pages

```
┌─────────────────────────────┐
│                             │
│       Auth Container        │
│       (Full viewport)       │
│                             │
│     ┌─────────────────────┐ │
│     │  Auth Card          │ │
│     │  (400px max width)  │ │
│     │                     │ │
│     │  Title              │ │
│     │  Subtitle           │ │
│     │                     │ │
│     │  ┌──────────────┐   │ │
│     │  │ Form Fields  │   │ │
│     │  └──────────────┘   │ │
│     │                     │ │
│     │  ┌──────────────┐   │ │
│     │  │ Submit Btn   │   │ │
│     │  └──────────────┘   │ │
│     │                     │ │
│     │  Toggle Link        │ │
│     └─────────────────────┘ │
│                             │
└─────────────────────────────┘
```

### Dashboard Page

```
┌────────────────────────────────────────────────┐
│  SecureVault          [Sign Out]               │
├────────────────────────────────────────────────┤
│                                                │
│  ┌──────────────────┐    ┌──────────────────┐ │
│  │ Upload Section   │    │ Files Section    │ │
│  │                  │    │                  │ │
│  │ Drag & drop area │    │ File List        │ │
│  │ or click         │    │ with actions     │ │
│  │                  │    │                  │ │
│  └──────────────────┘    └──────────────────┘ │
│                                                │
└────────────────────────────────────────────────┘
```

## Animations

### Transitions

```css
/* Default transition */
transition: all 0.2s ease;

/* Fade in */
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

/* Slide up */
@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Spin (loading) */
@keyframes spin {
  to { transform: rotate(360deg); }
}
```

### Interaction Feedback

- **Button Hover:** 2px upward movement + shadow increase
- **Button Active:** No movement, slight color shift
- **Input Focus:** 3px blue outline
- **Card Hover:** Slight shadow increase

## Responsive Design

### Breakpoints

```css
/* Mobile: 0px - 600px */
@media (max-width: 600px) {
  /* Single column layouts */
  /* Larger touch targets */
  /* Adjusted padding */
}

/* Tablet: 601px - 1024px */
@media (max-width: 1024px) {
  /* Two column layouts */
  /* Adjusted spacing */
}

/* Desktop: 1025px+ */
/* Full width layouts */
```

### Mobile Considerations

- Minimum touch target: 44x44px
- Padding adjusted for smaller screens
- Single column layouts
- Larger text sizes
- Simplified navigation

## Accessibility

### Color Contrast

- Normal text: Minimum 4.5:1 (AA standard)
- Large text: Minimum 3:1
- All interactive elements: Minimum 4.5:1

### Focus States

All interactive elements have visible focus states:

```css
:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(0, 102, 204, 0.1);
}
```

### Keyboard Navigation

- All buttons and inputs are keyboard accessible
- Tab order follows visual flow
- Enter key submits forms
- Escape key closes modals (when implemented)

### Screen Reader Support

- Semantic HTML elements
- Proper label associations
- ARIA attributes where needed
- Descriptive link text

## Future Enhancements

1. **Dark Theme** - High contrast dark mode option
2. **Animations** - Smooth page transitions
3. **Notifications** - Toast notifications
4. **File Preview** - Image/document preview
5. **Sharing** - File sharing with expiration
6. **Two-Factor Auth** - Enhanced security
7. **Advanced Search** - File search and filtering

---

**Design System Version:** 1.0  
**Last Updated:** May 2026  
**Created for:** GoSecureTransfer
