import { Card, Image, Text, Badge, Button, Group, CardSection, Accordion } from '@mantine/core';
import RepositoryType from '../interfaces/repository';
import Releases from './releases';

const Repository = ({ owner, name, releases }: RepositoryType) => {
    return (
        <Card shadow="sm" p="md" radius="md" withBorder>
            <CardSection >
                <Badge size='md' m='sm' variant='outline' radius="xs">
                    {owner}
                </Badge>
            </CardSection>
            <Group position="apart" mt="md" mb="xs">
                <Text weight={500}>{name}</Text>
            </Group>
            <Accordion>
                <Accordion.Item value="releases">
                    <Accordion.Control>Releases</Accordion.Control>
                    <Accordion.Panel>
                        {releases != null ? <Releases repoOwner={owner} repoName={name} releases={releases}></Releases> : <p>No releases</p>}
                    </Accordion.Panel>
                </Accordion.Item>
            </Accordion>

        </Card>)
}

export default Repository;



