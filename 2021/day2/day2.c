#include <stdio.h>

struct Position
{
    int Horizontal;
    int Depth;
    int DepthAim;
    int Aim;
};

int main()
{
    FILE *fp;
    struct Position pos;
    pos.Depth = 0;
    pos.DepthAim = 0;
    pos.Horizontal = 0;
    pos.Aim = 0;
    fp = fopen("input.txt", "r");
    char op[10];
    while (fscanf(fp, "%s", op) == 1)
    {
        int c;
        fscanf(fp, "%d", &c);

        if (op[0] == 'f')
        {
            pos.Horizontal += c;
            pos.DepthAim += pos.Aim * c;
        }
        else if (op[0] == 'd')
        {
            pos.Aim += c;
            pos.Depth += c;
        }
        else if (op[0] == 'u')
        {
            pos.Aim -= c;
            pos.Depth -= c;
        }
        else
        {
            return 1;
        }
    }

    fclose(fp);
    fprintf(stdout, "part 1: %d \n", pos.Depth * pos.Horizontal);
    fprintf(stdout, "part 2: %d \n", pos.DepthAim * pos.Horizontal);
}