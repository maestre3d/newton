import { Link } from 'react-router-dom'
import { APPLICATION_NAME } from '../../../../internal/shared/domain/newton'

const HorizontalLogo = (
    <Link to='/' className='flex flex-row items-center'>
        <span className='w-8 h-8 rounded-full bg-gray-300 mr-2' />
        <span className='font-bold'>{APPLICATION_NAME}</span>
    </Link>
)

export default HorizontalLogo