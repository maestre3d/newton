import { Fragment } from 'react'
import { Helmet } from 'react-helmet'
import { APPLICATION_NAME, APPLICATION_DOMAIN } from '../../../internal/shared/domain/newton'

function Explore() {
    return (
        <Fragment>
            <Helmet>
                <title>{`${APPLICATION_NAME} - Explore`}</title>
                <link rel='canonical' href={`https://${APPLICATION_DOMAIN}/explore`} />
            </Helmet>
            <div className='dark:text-gray-900'>Explore</div>
        </Fragment>
    )
}

export default Explore