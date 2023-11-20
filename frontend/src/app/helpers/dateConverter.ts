const dateOptions = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric',
    hour12: true
} as Intl.DateTimeFormatOptions
export default function convertDate(date: string) {
    return new Date(String(date).replace("UTC", "GMT")).toLocaleString("en-CA", dateOptions)
}
