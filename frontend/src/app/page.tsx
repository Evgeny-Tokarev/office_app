import dynamic from 'next/dynamic';
import Navbar from '@/components/Navbar';

const DashboardNoSSR = dynamic(
    () => import('@/app/dashboard/page'),
    { ssr: false }
)

export default function Page() {
    return (<>
        <Navbar/>
        <DashboardNoSSR/>
    </>)
}
