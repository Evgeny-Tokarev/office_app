import dynamic from 'next/dynamic';

const DashboardNoSSR = dynamic(
    () => import('@/app/dashboard/page'),
    { ssr: false }
)

export default function Page() {
    return (<>
        <DashboardNoSSR/>
    </>)
}
