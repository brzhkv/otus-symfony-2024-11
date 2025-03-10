<?php
declare(strict_types=1);

namespace App\Controller\Cli;

use App\Domain\Model\CreateUserModel;
use App\Domain\ValueObject\CommunicationChannelEnum;
use App\Domain\Service\SubscriptionService;
use App\Domain\Service\UserService;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Command\LockableTrait;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;
use Throwable;

/**
 * @author Mikhail Kamorin aka raptor_MVK
 *
 * @copyright 2020, raptor_MVK
 */
final class AddFollowersCommand extends Command
{
    use LockableTrait;

    /** @var int */
    public const OK = 0;
    /** @var int */
    public const GENERAL_ERROR = 1;

    public const DEFAULT_FOLLOWERS = 10;

    /** @var UserService */
    private $userService;
    /** @var SubscriptionService */
    private $subscriptionService;

    public function __construct(UserService $userService, SubscriptionService $subscriptionService)
    {
        parent::__construct();
        $this->userService = $userService;
        $this->subscriptionService = $subscriptionService;
    }

    protected function configure(): void
    {
        $this->setName('followers:add')
            ->setDescription('Adds followers to author')
            ->addArgument('authorId', InputArgument::REQUIRED, 'ID of author')
            ->addOption('count', 'c',  InputOption::VALUE_OPTIONAL, 'How many followers should be added', self::DEFAULT_FOLLOWERS)
            ->addOption('login', 'l',  InputOption::VALUE_REQUIRED, 'Follower Login');
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $authorId = (int)$input->getArgument('authorId');
        $user = $this->userService->findUserById($authorId);
        if ($user === null) {
            $output->write("<error>User with ID $authorId doesn't exist</error>\n");
            return self::GENERAL_ERROR;
        }
        $count = (int)($input->getOption('count') ?? self::DEFAULT_FOLLOWERS);
        if ($count < 0) {
            $output->write("<error>Count should be positive integer</error>\n");
            return self::GENERAL_ERROR;
        }
        $createdFollowers = 0;
        for ($i = 0; $i < $count; $i++) {
            try {
                $saveUserDTO = new CreateUserModel(
                    $login = "{$input->getOption('login')}.$i" ?? "Reader #$authorId.$i",
                    'no@mail.com',
                    CommunicationChannelEnum::Email,
                    "{$login}_password"
                );
                $follower = $this->userService->create($saveUserDTO);
                $this->subscriptionService->addSubscription($user, $follower);
                $createdFollowers++;
            } catch (Throwable $e) {
                $output->write("<error>User #$i couldn't be created</error>\n");
            }
        }
        $output->write("<info>$createdFollowers followers were created</info>\n");

        return self::OK;
    }
}