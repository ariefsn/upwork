import dayjs from "dayjs"
import type { MetaTagsProps } from "svelte-meta-tags"

export const formatDate = (d: Date | string, format?: string) => {
  return dayjs(d).format(format)
}

export const formatCurrency = (value: number) => {
  return new Intl.NumberFormat(
    "en-US",
    {
      currency: "USD",
      style: 'currency',
      minimumFractionDigits: 2,
      maximumFractionDigits: 2
    }
  ).format(value)
}

export const capitalize = (s: string) => {
  if (typeof s !== 'string') return ''
  return s.charAt(0).toUpperCase() + s.slice(1)
}

export const generatePageMetaData = ({ url, appName, title: _title, description, tags, image: _image }: { url: URL, appName: string, title: string, description: string, tags: string[], image?: string, }): MetaTagsProps => {
  const webUrl = new URL(url.pathname, url.origin).href;
  const defaultImage = 'https://images.unsplash.com/photo-1669399213378-2853e748f217?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'
  const image = _image || defaultImage
  const imageExt = image.split('.').pop()
  const title = [appName, _title].join(' | ')

  return {
    title,
    description,
    openGraph: {
      type: 'article',
      title,
      description,
      siteName: appName,
      url: webUrl,
      article: {
        authors: ['Arief Setiyo Nugroho'],
        tags,
        publishedTime: new Date().toISOString(),
      },
      images: [
        {
          url: image,
          alt: description || title,
          width: 800,
          height: 600,
          secureUrl: image,
          type: 'image/' + imageExt
        }
      ]
    },
    keywords: tags,
    twitter: {
      site: '@ariefsn04',
      title,
      cardType: 'summary_large_image',
      description,
      image: image,
      imageAlt: description || title
    },
  }
}