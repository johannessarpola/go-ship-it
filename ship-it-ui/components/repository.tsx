import { Card, Image, Text, Badge, Button, Group, CardSection } from '@mantine/core';
import RepositoryType from '../interfaces/repository';
import Releases from './releases';

const Repository = ({ name, releases }: RepositoryType) => {
    const abc = 123;
    // TODO Add org avatar
    return (
        <Card shadow="sm" p="lg" radius="md" withBorder>
            <CardSection >
                <Badge size='lg' m='sm' variant='outline' radius="xs">
                    {name}
                </Badge>
            </CardSection>
            <Group position="apart" mt="md" mb="xs">
                { releases != null ?  <Releases releases={releases}></Releases> : "" }
                
            </Group>
        </Card>)
}

export default Repository;



