<?php

namespace FunctionalTests\Controller\Cli;

use App\Controller\Cli\AddFollowersCommand;
use App\Domain\Entity\EmailUser;
use App\Domain\Entity\PhoneUser;
use App\Tests\Support\FunctionalTester;
use Codeception\Example;

class AddFollowersCommandCest
{
    private const COMMAND = 'followers:add';

    /**
     * @dataProvider executeDataProvider
     */
    public function testExecuteReturnsResult(FunctionalTester $I, Example $example): void
    {
        /** @var PhoneUser $author */
        $author = $I->have(PhoneUser::class);

        $params = ['authorId' => $author->getId(), '--login' => $example['login']];
        if (null !== $example['followersCount']) {
            $params['--count'] = $example['followersCount'];
        }
        $inputs = ["\n"];
        $output = $I->runSymfonyConsoleCommand(self::COMMAND, $params, $inputs, $example['exitCode'] ?? 0);
        $I->assertStringEndsWith($example['expected'], $output);
    }

    protected function executeDataProvider(): array
    {
        return [
            'positive' => ['followersCount' => 20, 'login' => 'login', 'expected' => "20 followers were created\n"],
            'zero' => ['followersCount' => 0, 'login' => 'other_login', 'expected' => "0 followers were created\n"],
            'default' => ['followersCount' => null, 'login' => 'login3', 'expected' => "10 followers were created\n"],
            'negative' => ['followersCount' => -1, 'login' => 'login_too', 'expected' => "Count should be positive integer\n", 'exitCode' => AddFollowersCommand::GENERAL_ERROR],
        ];
    }
}